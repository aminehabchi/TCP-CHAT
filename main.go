package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"funcs/funcs"
)

func main() {
	port := ":8989"
	if len(os.Args) == 2 {
		n, e := strconv.Atoi(os.Args[1])
		if e == nil && n >= 0 && n <= 65535 {
			port = ":" + os.Args[1]
		} else {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}

	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	// turncate the prev message file
	file, err := os.OpenFile("./Logs&&PreMessage/prevMessages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	file.Truncate(0)

	///save start of new chat with time
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	funcs.SaveToFile("./Logs&&PreMessage/logs.txt", "------------------------new chat started at ["+formattedTime+"]--------------------------------\n\n\n")

	// Start TCP server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server started on port ", port)

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go funcs.HandleClient(conn)
	}
}
