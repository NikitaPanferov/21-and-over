package main

import (
	"fmt"
	"log"

	tcpclient "github.com/NikitaPanferov/21-and-over/client/pkg/tcp-client"
)

func main() {
	client, err := tcpclient.NewClient("192.168.1.75:9000")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	// Отправляем сообщение типа ECHO.
	response, err := client.SendMessage("ECHO", []byte("Hello, server!"))
	if err != nil {
		log.Fatalf("Error sending ECHO message: %v", err)
	}
	fmt.Printf("ECHO response: %s\n", string(response))

	// Отправляем сообщение типа ACK.
	response, err = client.SendMessage("ACK", []byte("Testing ACK handler"))
	if err != nil {
		log.Fatalf("Error sending ACK message: %v", err)
	}
	fmt.Printf("ACK response: %s\n", string(response))
}
