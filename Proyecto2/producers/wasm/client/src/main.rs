use std::net::SocketAddr;
use hyper::server::conn::Http;
use hyper::service::service_fn;
use hyper::{Body, Method, Request, Response, StatusCode, Client}; // Added Client here
use tokio::net::TcpListener;



// Function to send the body to the server that dispatch the message to Kafka
async fn send_to_server(body: Body) -> Result<Response<Body>, hyper::Error> {
    // Convert the body to bytes.
    let body_bytes = hyper::body::to_bytes(body).await?;


    // Create a new Client.
    let client = Client::new();

    // Create a new Request.
    let req = Request::builder()
    .method(Method::POST)
    .uri("http://0.0.0.0:8081/wasm")
    .header("Content-Type", "application/json") // Add this line
    .body(Body::from(body_bytes))
    .expect("request builder");

    // Send the request and await the response.
    let res = client.request(req).await?;

    println!("Response: {:?}", res);

    Ok(res)
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
        
            match send_to_server(body).await {
                Ok(_res) => Ok(Response::new(Body::from("Message sent to other service"))),
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