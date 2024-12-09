package main

import (
	"fmt"

	controllerPkg "github.com/NikitaPanferov/21-and-over/server/internal/controller"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func main() {
	server := tcpserver.NewServer("localhost:9000")

	gameService := struct{}{}
	controller := controllerPkg.New(gameService)

	controllerPkg.RegisterHandlers(server, controller)

	// Запускаем сервер.
	if err := server.Start(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
