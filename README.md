# TCP Chat Server

A TCP chat server built in both **Rust** and **Go** that creates a group chat environment. Users connect to the server using **netcat** (`nc`) or any TCP client to participate in real-time group messaging.

## 🎯 Objectives

This project creates a TCP server that recreates NetCat's chat functionality:
- Runs a TCP server on a specified port, listening for incoming connections
- Allows multiple users to connect simultaneously using netcat (`nc` command)
- Provides a group chat environment where all connected users can communicate
- Manages client connections and broadcasts messages between all participants

## ✨ Features

### Core Functionality
- **TCP Connection**: Establishes TCP connections between server and multiple clients (1-to-many relationship)
- **User Authentication**: Requires clients to provide a name when joining
- **Connection Management**: Controls and monitors connection quantities
- **Real-time Messaging**: Clients can send and receive messages instantly
- **Message Broadcasting**: All connected clients receive messages from other participants

### Message Features
- **Empty Message Filtering**: Prevents broadcasting of empty messages
- **Message Formatting**: All messages include timestamp and sender information
  ```
  [2020-01-20 15:48:41][client.name]:[client.message]
  ```
- **Message History**: New clients receive all previous chat messages upon joining

### Connection Events
- **Join Notifications**: Server notifies all clients when someone joins the chat
- **Leave Notifications**: Server notifies all clients when someone leaves the chat
- **Persistent Connections**: Remaining clients stay connected when others leave

### Configuration
- **Default Port**: Uses port `8989` when no port is specified
- **Custom Port**: Accepts custom port via command line argument
- **Usage Validation**: Shows usage message for incorrect arguments

