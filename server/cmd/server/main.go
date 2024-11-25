package main

import (
	"fmt"

	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func echoHandler(ctx *tcpserver.Context) error {
	fmt.Printf("Echo handler received: %s\n", string(ctx.GetMessage()))
	return ctx.Write(ctx.GetMessage())
}

func ackHandler(ctx *tcpserver.Context) error {
	fmt.Printf("ACK handler received: %s\n", string(ctx.GetMessage()))
	return ctx.Write([]byte("ACK"))
}

func main() {
	server := tcpserver.NewServer("192.168.1.75:9000")

	// Регистрируем обработчики.
	server.RegisterHandler("ECHO", echoHandler)
	server.RegisterHandler("ACK", ackHandler)

	// Запускаем сервер.
	if err := server.Start(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
