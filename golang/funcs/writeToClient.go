package funcs

import "time"

func SendToAntherClients(message, clientAddr string) {
	for i, j := range clients {
		if i != clientAddr {
			if j.Name != "" {
				j.Conn.Write([]byte(message))
			}
		}
	}
}

func WriteToClients(message string, clientAddr string, bl bool) {
	if bl {
		message = "\n" + GeneateMessage(clients[clientAddr].Name) + message
		SaveToFile("./Logs&&PreMessage/prevMessages.txt", message[1:])
	} else {
		message = "\n" + clients[clientAddr].Name + message
		SaveToFile("./Logs&&PreMessage/logs.txt", GeneateMessage("Client Name: "+clients[clientAddr].Name+" || Client Adress "+clientAddr)+message[1:])
	}

	SendToAntherClients(message, clientAddr)
}

func GeneateMessage(name string) string {
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	return "[" + formattedTime + "]" + "[" + name + "]" + ":"
}

func Status() {
	for i, j := range clients {
		if j.Name != "" {
			j.Conn.Write([]byte(GeneateMessage(clients[i].Name)))
		}
	}
}
