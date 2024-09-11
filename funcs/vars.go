package funcs

import (
	"net"
	"sync"
)

type Client struct {
	Name string
	Conn net.Conn
}

var (
	clients      = make(map[string]Client) // Map to store connected clients
	clientsMutex sync.Mutex                // Mutex to synchronize access to clients map
)

const MAX_CLIENTS int = 4
