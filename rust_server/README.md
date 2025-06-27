
# Net-Cat: Simple TCP Chat Server in Rust

Net-Cat is a basic multi-client TCP chat server written in Rust. It allows multiple clients to connect, enter their names, and exchange messages in real-time. Messages are broadcast to all connected clients with a timestamp and the sender's name.

---

## Features

- Multi-client support using threads and shared state
- Client name registration upon connection
- Broadcast messages with timestamps and sender's username
- Validation to reject messages containing non-printable ASCII characters
- Graceful handling of client disconnections

---

## Getting Started

### Prerequisites

- Rust (latest stable version recommended)
- Cargo (comes with Rust)

### Build and Run

Clone the repository and run:

```bash
cargo run
```

This starts the server listening on `127.0.0.1:8080`.

### Connect as a client

You can use tools like `telnet` or `netcat` to connect:

```bash
telnet 127.0.0.1 8080
```

or

```bash
nc 127.0.0.1 8080
```

Enter your username when prompted, then start chatting!

---

## Code Structure

- `main.rs`: Starts the server, listens for incoming connections, and spawns client handler threads.
- `clients.rs`: Defines the `Client` struct and manages client name registration.
- `talking.rs`: Contains logic for broadcasting messages, handling clients, and removing disconnected clients.

---

## Notes

- This project does **not** use asynchronous programming libraries like Tokio; it uses threads and `Arc<Mutex<>>` for concurrency.
- Message validation ensures only printable ASCII characters (except newline) are broadcasted.
- Feel free to extend or modify the server for more features!

---

## License

MIT License

---

Enjoy chatting!
