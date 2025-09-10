package controller

import (
	"errors"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func (c *Controller) standHandler(ctx *tcpserver.Context) error {
	allowed := c.ValidateState(ctx)
	if !allowed {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			errors.New("invalid state"),
		)
	}

	gameState, err := c.gameService.Stand(ctx.GetContext(), ctx.GetSender())
	if err != nil {
		switch {
		case !errors.Is(err, entities.ErrPlayerShouldBeAlreadyDone):
			break
		default:
			return ctx.ReplyWithError(
				tcpserver.CodeClientError,
				fmt.Errorf("c.gameService.Hit: %w", err),
			)
		}
	}

	ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
		EventType: tcpserver.EventTypePlayerDidStand,
		Data:      gameState,
	})

	if c.gameService.GetState(ctx.GetContext()) == entities.ResultState {
		ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
			EventType: tcpserver.EventTypeGameEnded,
			Data:      gameState,
		})

		c.gameService.Reset(ctx.GetContext())
		ctx.ClearConnections()
	}

	return nil
}
