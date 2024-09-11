package funcs

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Function to handle each client connection
func HandleClient(conn net.Conn) {
	defer conn.Close()

	// check number of clints
	clientsMutex.Lock()
	if len(clients) >= MAX_CLIENTS {
		clientsMutex.Unlock()
		conn.Write([]byte("Connection is Full in the server"))
		return
	}
	clientsMutex.Unlock()

	// get ip addres of clients
	clientAddr := conn.RemoteAddr().String()

	// penguin
	penguin, e := os.ReadFile("Logs&&PreMessage/penguin.txt")
	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}
	conn.Write(penguin)

	// get and check name
	for !IsValidName(clients[clientAddr].Name) {
		// read name
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		reader := bufio.NewReader(conn)
		name, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		name = name[:len(name)-1]
		clientsMutex.Lock()
		clients[clientAddr] = Client{Conn: conn, Name: name}
		clientsMutex.Unlock()
	}

	// read message
	reader := bufio.NewReader(conn)

	// write to anthors clients enter message
	WriteToClients(" has joined our chat...\n", clientAddr, false)

	// write to the new client prevMessage
	clients[clientAddr].Conn.Write([]byte(prevMessage()))

	bl := true
	for {
		// check for empty message
		if bl {
			Status()
		}
		bl = true
		// Read data until newline or EOF (you can modify the delimiter if needed)

		message, err := reader.ReadString('\n')
		if err != nil {
			// leave message
			WriteToClients(" has left our chat...\n", clientAddr, false)
			Status()
			delete(clients, clientAddr)
			break // Exit loop if the client disconnects or an error occurs
		}

		// send message to clients
		if !isValidMessage(message) {
			conn.Write([]byte(GeneateMessage(clients[clientAddr].Name)))
			bl = false
		} else {
			// write to anthors clients the message
			WriteToClients(message, clientAddr, true)
		}
	}

	// Remove client from the map when they disconnect
	delete(clients, clientAddr)
}
