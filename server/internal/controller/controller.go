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
	server.RegisterHandler("JOIN", controller.joinHandler)
}

func (c *Controller) echoHandler(ctx *tcpserver.Context) error {
	err := ctx.SendToAll(tcpserver.CodeSuccess, ctx.GetMessage().Data.(string))
	if err != nil {
		return fmt.Errorf("ctx.SendToAll: %w", err)
	}

	return nil
}
