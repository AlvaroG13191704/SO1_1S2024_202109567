
use std::net::SocketAddr;
use hyper::server::conn::Http;
use hyper::service::service_fn;
use hyper::{Body, Method, Request, Response, StatusCode};
use tokio::net::TcpListener;
use rskafka::{
    client::{
        ClientBuilder,
        partition::{Compression, UnknownTopicHandling},
    },
    record::Record,
};
use chrono::Utc;
use std::collections::BTreeMap;
use serde_json::Value;

// produce a event to kafka
async fn produce(event: Value) -> Result<(), Box<dyn std::error::Error>> {
    let brokers = "my-cluster-kafka-bootstrap:9092";
    let event_str = event.to_string();

    let client = ClientBuilder::new(vec![brokers.to_owned()]).build().await?;

    let topic = "mytopic";
    let controller_client = client.controller_client().unwrap();
    controller_client.create_topic(
        topic,
        0,      // partitions
        1,      // replication factor
        5_000,  // timeout (ms)
    ).await?;

    let partition_client = client
        .partition_client(
            topic.to_owned(),
            0,  // partition
            UnknownTopicHandling::Retry,
        )
        .await?;

        let record = Record {
            key: None,
            value: Some(event_str.into_bytes()),
            headers: BTreeMap::new(),
            timestamp: Utc::now(),
        };
    partition_client.produce(vec![record], Compression::default()).await?;

    Ok(())
}
/// This is our service handler. It receives a Request, routes on its
/// path, and returns a Future of a Response.
async fn echo(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    match (req.method(), req.uri().path()) {
        // Serve some instructions at /
        (&Method::GET, "/") => Ok(Response::new(Body::from(
            "Try PoSTing data to /echo such as: `curl localhost:3000/echo -XPOST -d 'hello world'`",
        ))),

        // Simply echo the body back to the client.
        (&Method::POST, "/wasm") => {
            let whole_body = hyper::body::to_bytes(req.into_body()).await?;
            let event: Value = serde_json::from_slice(&whole_body).unwrap();
            match produce(event).await {
                Ok(_) => Ok(Response::new(Body::from("Message sent to Kafka"))),
                Err(e) => {
                    let mut response = Response::new(Body::from(format!("Failed to send message: {}", e)));
                    *response.status_mut() = StatusCode::INTERNAL_SERVER_ERROR;
                    Ok(response)
                }
            }
        }

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
    let addr = SocketAddr::from(([0, 0, 0, 0], 3000));

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