package controller

import (
	"fmt"

	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

type (
	GameService interface{}

	Controller struct {
		gameService GameService
	}
)

func New(gameService GameService) *Controller {
	return &Controller{
		gameService: gameService,
	}
}

func RegisterHandlers(server *tcpserver.Server, controller *Controller) {
	server.RegisterHandler("ECHO", controller.echoHandler)
	server.RegisterHandler("ACK", controller.ackHandler)
}

func (c *Controller) echoHandler(ctx *tcpserver.Context) error {
	fmt.Printf("Echo handler received: %s\n", string(ctx.GetMessage()))
	return ctx.Write(ctx.GetMessage())
}

func (c *Controller) ackHandler(ctx *tcpserver.Context) error {
	fmt.Printf("ACK handler received: %s\n", string(ctx.GetMessage()))
	return ctx.Write([]byte("ACK"))
}
