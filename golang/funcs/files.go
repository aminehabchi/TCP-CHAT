package funcs

import (
	"fmt"
	"os"
)

func SaveToFile(name, message string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	file.WriteString(message)
}

func prevMessage() string {
	data, err := os.ReadFile("./Logs&&PreMessage/prevMessages.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return string(data)
}
