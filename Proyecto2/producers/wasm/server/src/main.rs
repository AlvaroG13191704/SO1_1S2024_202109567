use axum::{
    routing::{get, post}, Json, Router
};

use serde::{Serialize, Deserialize};
use serde_json;
use std::convert::Infallible;
// use rdkafka::message::ToBytes;
use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::util::Timeout;

#[derive(Serialize, Deserialize, Debug)]
struct AlbumData {
    name: String,
    album: String,
    year: String,
    rank: String,
}

// impl ToBytes for AlbumData {
//     fn to_bytes(&mut self) -> &[u8] {
//         // Serialize self to a JSON string
//         let json = serde_json::to_string(self).unwrap();

//         // Convert the JSON string to bytes and store it in the struct
//         self.bytes = Some(json.into_bytes());

//         // Return a reference to the bytes
//         self.bytes.as_ref().unwrap().as_slice()
//     }
// }

#[tokio::main]
async fn main() {
    // build our application with a single route
    let app = Router::new()
    .route("/", get(|| async { "Hello, World!" }))
    .route("/wasm", post(get_album));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8081").await.unwrap();
    println!("Server running on port 8081");
    axum::serve(listener, app).await.unwrap();
}

// Function to received the post request
async fn get_album(Json(payload) : Json<AlbumData>) -> Result<Json<String>, Infallible> {
    // print
    println!("Received: {:?}", payload);

    // convert to string
    let payload_string = match serde_json::to_string(&payload) {
        Ok(v) => v,
        Err(e) => {
            println!("Failed to serialize payload to JSON: {:?}", e);
            "Error".to_string() // return an empty string if there's an error
        }
    };

    // create producer
    let producer: FutureProducer = rdkafka::ClientConfig::new()
        .set("bootstrap.servers", "my-cluster-kafka-bootstrap:9092")
        .create()
        .expect("Producer creation error");


    // Produce a message to Kafka
    let topic = "mytopic";
    let delivery_status = producer.send(
        FutureRecord::<String, _>::to(&topic)
            .payload(&payload_string.as_bytes().to_vec()),
        Timeout::Never,
    ).await;
    
    match delivery_status {
        Ok((partition, offset)) => println!("Message delivered to partition {} at offset {}", partition, offset),
        Err((e, _)) => println!("Error producing message: {:?}", e),
    }

    // return the payload
    Ok(Json("Received".to_string()))
}