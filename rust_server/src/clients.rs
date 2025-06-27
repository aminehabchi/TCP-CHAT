use std::{
    io::{Read, Write},
    net::TcpStream,
};

pub struct Client {
    pub stream: TcpStream,
    pub name: String,
}

impl Client {
    pub fn new(stream: TcpStream, name: String) -> Self {
        Client { stream, name }
    }

    pub fn get_name(&mut self) {
           let logo = "Welcome to TCP-Chat!
_ _nnnn_
dGGGGMMb
@p~qp~~qMb
M|@||@) M|
@,----.JM|
JS^\\__/  qKL
dZP        qKRb
dZP          qKKb
fZP            SMMb
HZM            MMMM
FqM            MMMM
__| \".        |\\dS\"qML
|    `.       | `' \\Zq
_)      \\.___.,|     .'
\\____   )MMMMMP|   .'
`-'       `--'\nPlease enter your name: ";

        let _ = self.stream.write_all(logo.as_bytes());

        let mut buffer = [0u8; 512];

        loop {
            match self.stream.read(&mut buffer) {
                Ok(0) => break, // Connection closed
                Ok(n) => {
                    let name = String::from_utf8_lossy(&buffer[..n]).trim().to_string();
                    if is_valid_name(&name) {
                        self.name = name;
                        break;
                    } else {
                        let _ = self.stream.write_all(b"Invalid name. Please enter your name again:\n");
                    }
                }
                Err(_) => break,
            }
        }
    }

    pub fn peer_addr(&self) -> std::io::Result<std::net::SocketAddr> {
        self.stream.peer_addr()
    }
}

// Implement manual Clone because TcpStream doesn't implement Clone
impl Clone for Client {
    fn clone(&self) -> Self {
        Client {
            stream: self.stream.try_clone().expect("Failed to clone TcpStream"),
            name: self.name.clone(),
        }
    }
}

fn is_valid_name(name: &str) -> bool {
    if name.trim().is_empty() {
        return false;
    }

    name.chars().all(|c| c.is_alphanumeric() || c == ' ')
}
