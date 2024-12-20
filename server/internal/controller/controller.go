package controller

import (
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

type (
	GameService interface{
		Join(ctx *tcpserver.Context, player entities.Player) error
	}

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
	server.RegisterHandler("JOIN", controller.joinHandler)
}

func (c *Controller) echoHandler(ctx *tcpserver.Context) error {
	fmt.Printf("Echo handler received: %s\n", string(ctx.GetMessage()))
	_ = ctx.SendToAll(ctx.GetMessage())
	fmt.Printf("Echo handler sent: %s\n", string(ctx.GetMessage()))
	return nil
}

func (c *Controller) ackHandler(ctx *tcpserver.Context) error {
	fmt.Printf("ACK handler received: %s\n", string(ctx.GetMessage()))
	return ctx.Write([]byte("ACK"))
}
