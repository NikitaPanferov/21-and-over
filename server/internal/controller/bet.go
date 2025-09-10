package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/controller/types"
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	"github.com/NikitaPanferov/21-and-over/server/pkg/logger"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func (c *Controller) betHandler(ctx *tcpserver.Context) error {
	allowed := c.ValidateState(ctx)
	if !allowed {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			errors.New("invalid state"),
		)
	}

	var request types.BetRequest
	err := json.Unmarshal(ctx.GetRawData(), &request)
	if err != nil {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			fmt.Errorf("json.Unmarshal: %w", err),
		)
	}

	gameState, err := c.gameService.Bet(ctx.GetContext(), ctx.GetSender(), request.Bet)
	if err != nil {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			fmt.Errorf("c.gameService.Bet: %w", err),
		)
	}

	err = ctx.Reply(tcpserver.CodeSuccess, nil)
	if err != nil {
		logger.ErrorContext(ctx.GetContext(), "ctx.Reply: %v", err)
	}

	ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
		EventType: tcpserver.EventTypePlayerDidBet,
		Data:      gameState,
	})

	if c.gameService.GetState(ctx.GetContext()) == entities.PlayState {
		ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
			EventType: tcpserver.EventTypeGameStarted,
			Data:      gameState,
		})
	}

	return nil
}
