// Ommitting Kafka initialisation boilerplate...

pub async fn receive_messages(cfg: ServiceConfig) {
    let consumer = create_consumer(...);
    let mut msg_stream = consumer.start();

    // iterate over all messages blocking
    while let Some(msg) = msg_stream.next().await {
        match msg {
            Ok(msg) => {
                // Process the message here...
            }
            Err(e) => error!("Could not receive and will not process message: {}", e),
        };
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    //...
    // Spawning the Kafka processor to run in parallel with the API 
    actix_web::rt::spawn(async move { receive_messages(config).await });
    //...
}