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

func (c *Controller) joinHandler(ctx *tcpserver.Context) error {
	allowed := c.ValidateState(ctx)
	if !allowed {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			errors.New("invalid state"),
		)
	}

	var request types.JoinRequest
	err := json.Unmarshal(ctx.GetRawData(), &request)
	if err != nil {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			fmt.Errorf("json.Unmarshal: %w", err),
		)
	}

	player := c.gameService.GetPlayer(request.Name, ctx.GetSender())

	gameState, err := c.gameService.Join(ctx.GetContext(), player)
	if err != nil {
		return ctx.ReplyWithError(
			tcpserver.CodeClientError,
			fmt.Errorf("c.gameService.Join: %w", err),
		)
	}

	err = ctx.Reply(tcpserver.CodeSuccess, nil)
	if err != nil {
		logger.ErrorContext(ctx.GetContext(), "ctx.Reply: %v", err)
	}

	ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
		EventType: tcpserver.EventTypePlayerJoined,
		Data:      gameState,
	})

	if c.gameService.GetState() == entities.WaitBetState {
		ctx.SendToAll(tcpserver.CodeSuccess, &tcpserver.Message{
			EventType: tcpserver.EventTypeWaitingBet,
			Data:      gameState,
		})
	}

	return nil
}
