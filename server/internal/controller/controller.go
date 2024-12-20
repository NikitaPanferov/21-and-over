package controller

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

type (
	GameService interface {
		Join(ctx context.Context, player *entities.Player) (*entities.GameState, error)
		GetState() entities.State
		GetPlayer(
			name, ip string,
		) *entities.Player
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
	server.RegisterHandler("JOIN", controller.joinHandler)
}

func (c *Controller) echoHandler(ctx *tcpserver.Context) error {
	ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
		ID:     ctx.GetMessage().ID,
		Action: ctx.GetMessage().Action,
		Data:   ctx.GetMessage().Data,
		Code:   tcpserver.CodeSuccess,
	})

	return nil
}
