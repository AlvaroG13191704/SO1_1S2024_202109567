use std::net::SocketAddr;

use hyper::server::conn::Http;
use hyper::service::service_fn;
use hyper::{Body, Method, Request, Response, StatusCode};
use tokio::net::TcpListener;
// kafka
use rskafka::{
    client::{
        ClientBuilder,
        partition::{Compression, UnknownTopicHandling},
    },
    record::Record,
};
use chrono::{TimeZone, Utc,LocalResult};
use std::collections::BTreeMap;

// Function to produce a message to kafka
async fn produce(event : Body)  -> Result<(), Box<dyn std::error::Error>> {
    println!("Producing message to Kafka");

    let body = hyper::body::to_bytes(event).await?;
    let event_str = String::from_utf8(body.to_vec()).unwrap();

    println!("step 1");
    let connection = "my-cluster-kafka-boostrap:9092".to_owned();
    let client = ClientBuilder::new(vec![connection]).build().await.map_err(|e| Box::new(e))?;

    println!("step 2");
    let topic = "mytopic";
    let controller_client = client.controller_client().unwrap();
    match controller_client.create_topic(topic, 2, 1, 5_000).await {
        Ok(_) => (),
        Err(e) => {
            println!("Error creating topic: {}", e);
            return Err(Box::new(e));
        }
    };

    println!("step 3");

    let partition_client = match client.partition_client(topic.to_owned(), 0, UnknownTopicHandling::Retry).await {
        Ok(client) => client,
        Err(e) => {
            println!("Error creating partition client: {}", e);
            return Err(Box::new(e));
        }
    };

    println!("step 4");

    let timestamp = match Utc.timestamp_millis_opt(42) {
        LocalResult::Single(dt) => dt,
        LocalResult::Ambiguous(dt1, _dt2) => dt1,
        LocalResult::None => Utc::now(),
    };


    println!("step 5");

    let record = Record {
        key: None,
        value: Some(event_str.into_bytes()),
        headers: BTreeMap::new(),
        timestamp: timestamp,
    };

    println!("step 6");
    
    match partition_client.produce(vec![record], Compression::default()).await {
        Ok(_) => (),
        Err(e) => {
            println!("Error producing message: {}", e);
            return Err(Box::new(e));
        }
    };

    println!("Message sent to Kafka");

    Ok(())
}

/// This is our service handler. It receives a Request, routes on its
/// path, and returns a Future of a Response.
async fn echo(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    match (req.method(), req.uri().path()) {
        // Serve some instructions at /
        (&Method::GET, "/") => Ok(Response::new(Body::from(
            "Try POSTing data to /echo such as: `curl localhost:8080/echo -XPOST -d 'hello world'`",
        ))),

        // Simply echo the body back to the client.
        (&Method::POST, "/wasm") => {
            let body = req.into_body();
        
            // print
            println!("Received a POST request");
        
            match produce(body).await {
                Ok(_body) => Ok(Response::new(Body::from("Message sent to Kafka"))),
                Err(e) => {
                    let error_message = format!("Error: {}", e);
                    Ok(Response::builder()
                        .status(StatusCode::INTERNAL_SERVER_ERROR)
                        .body(Body::from(error_message))
                        .unwrap())
                },
            }
        },

        // Return the 404 Not Found for other routes.
        _ => {
            let mut not_found = Response::default();
            *not_found.status_mut() = StatusCode::NOT_FOUND;
            Ok(not_found)
        }
    }
}

#[tokio::main(flavor = "current_thread")]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let addr = SocketAddr::from(([0, 0, 0, 0], 8080));

    let listener = TcpListener::bind(addr).await?;
    println!("Listening on http://{}", addr);
    loop {
        let (stream, _) = listener.accept().await?;

        tokio::task::spawn(async move {
            if let Err(err) = Http::new().serve_connection(stream, service_fn(echo)).await {
                println!("Error serving connection: {:?}", err);
            }
        });
    }
}