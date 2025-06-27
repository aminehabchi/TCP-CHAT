use std::{
    io::Result,
    net::TcpListener,
    sync::{Arc, Mutex},
    thread,
};

mod clients;
use clients::*;

mod talking;
use talking::*;

fn main() -> Result<()> {
    let listener = TcpListener::bind("127.0.0.1:8080")?;
    println!("Server listening on 127.0.0.1:8080");

    let all_clients = Arc::new(Mutex::new(Vec::<Client>::new()));

    for stream_result in listener.incoming() {
        match stream_result {
            Ok(stream) => {
                // Clone Arc for thread
                let all_clients = Arc::clone(&all_clients);

                thread::spawn(move || {
                    // Create client here with empty name initially
                    let mut client = Client::new(stream.try_clone().unwrap(), String::new());

                    // Get client name here — blocks only this thread
                    client.get_name();

                    // Now add client to shared list
                    {
                        let mut clients = all_clients.lock().unwrap();
                        clients.push(client.clone());
                    }

                    // Start handling communication
                    handle_client(client, &all_clients);
                });
            }
            Err(e) => {
                eprintln!("Connection failed: {}", e);
            }
        }
    }

    Ok(())
}
