use crate::Client;
use std::{
    io::{Read, Write},
    net::TcpStream,
    sync::{Arc, Mutex},
};

use chrono::Local; 
pub fn broadcast_message(all_clients: &Arc<Mutex<Vec<Client>>>, sender: &TcpStream, raw_message: &[u8]) {
    if raw_message.len() > 1 && raw_message[..raw_message.len() - 1].iter().any(|&b| b < 32 || b > 126) {
        eprintln!("Message contains non-printable ASCII characters (excluding last char), skipping broadcast.");
        return;
    }

    let mut clients = all_clients.lock().unwrap();
    let sender_addr = sender.peer_addr().ok();

    let sender_name = clients
        .iter()
        .find(|c| c.peer_addr().ok() == sender_addr)
        .map(|c| c.name.clone())
        .unwrap_or_else(|| "Unknown".to_string());


    let message_text = String::from_utf8_lossy(raw_message).trim().to_string();

    if message_text.is_empty() {
        return;
    }

    let timestamp = Local::now().format("%Y-%m-%d %H:%M:%S").to_string();

    let formatted_msg = format!("[{}][{}]: {}\n", timestamp, sender_name, message_text);

    for client in clients.iter_mut() {
        if client.peer_addr().ok() != sender_addr {
            if let Err(e) = client.stream.write_all(formatted_msg.as_bytes()) {
                eprintln!("Failed to write to a client: {}", e);
            }
        }
    }
}


pub fn delete_client(all_clients: &Arc<Mutex<Vec<Client>>>, sender: &TcpStream) {
    let sender_addr = sender.peer_addr().ok();
    let mut clients = all_clients.lock().unwrap();

    if let Some(pos) = clients.iter().position(|c| c.peer_addr().ok() == sender_addr) {
        clients.remove(pos);
    }
}

pub fn handle_client(mut client: Client, all_clients: &Arc<Mutex<Vec<Client>>>) {
    let mut buffer = [0u8; 512];

    loop {
        match client.stream.read(&mut buffer) {
            Ok(0) => {
                println!("{} disconnected", client.name);
                delete_client(all_clients, &client.stream);
                break;
            }
            Ok(n) => {
                broadcast_message(all_clients, &client.stream, &buffer[..n]);
            }
            Err(e) => {
                eprintln!("Failed to read from {}: {}", client.name, e);
                delete_client(all_clients, &client.stream);
                break;
            }
        }
    }
}
