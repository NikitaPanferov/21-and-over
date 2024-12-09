package main

import (
	"encoding/json"
	"fmt"
	"log"

	tcpclient "github.com/NikitaPanferov/21-and-over/client/pkg/tcp-client"
)

func main() {
	client, err := tcpclient.NewClient("localhost:9000")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	for i := 1; i < 10; i++ {
		// Отправляем сообщение типа ECHO.
		dataToSend, err := json.Marshal(map[string]string{
			"name": fmt.Sprintf("Player %d", i),
		})
		response, err := client.SendMessage("JOIN", dataToSend)
		if err != nil {
			log.Fatalf("Error sending JOIN message: %v", err)
		}
		fmt.Printf("JOIN response: %s\n", response)

		dataFromServer := struct {
			Code uint16 `json:"code"`
			Data any    `json:"data"`
		}{}
		err = json.Unmarshal(response, &dataFromServer)
		if err != nil {
			log.Fatalf("Error unmarshalling JOIN response: %v", err)
		}

		fmt.Printf("JOIN response: %v\n", dataFromServer)

		fmt.Scanf("\n")
	}
}
