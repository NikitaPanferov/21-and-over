package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/controller/types"
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func (c *Controller) joinHandler(ctx *tcpserver.Context) error {
	fmt.Printf("Join handler received: %s\n", string(ctx.GetMessage()))

	var request types.JoinRequest
	err := json.Unmarshal(ctx.GetMessage(), &request)
	if err != nil {
		writeErr := ctx.Write(tcpserver.Response{
			Code: tcpserver.ResponseCodeClientError,
			Data: map[string]string{
				"error": err.Error(),
			},
		})
		if writeErr != nil {
			return fmt.Errorf("ctx.Write: %w", err)
		}

		//TODO: сделать норм логгер
		fmt.Printf("INFO: handled error: %v\n", err)

		return nil
	}

	player := entities.NewPlayer(request.Name)

	err = c.gameService.Join(ctx, *player)
	if err != nil {
		switch {
		case errors.Is(err, entities.ErrJoinGameIsFull):
		case errors.Is(err, entities.ErrJoinGameIsOn):
		case errors.Is(err, entities.ErrJoinPlayerAlreadyInGame):
			writeErr := ctx.Write(tcpserver.Response{
				Code: tcpserver.ResponseCodeClientError,
				Data: map[string]string{
					"error": err.Error(),
				},
			})
			if writeErr != nil {
				return fmt.Errorf("ctx.Write: %w", err)
			}
			return nil
		default:
			return fmt.Errorf("c.gameService.Join: %w", err)
		}
	}

	err = ctx.Write(tcpserver.Response{
		Code: tcpserver.ResponseCodeSuccess,
	})
	if err != nil {
		return fmt.Errorf("ctx.Write: %w", err)
	}

	ctx.WriteAll(c.gameService.)

	//TODO: сделать норм логгер
	fmt.Printf("INFO: handled join request for: %s\n", ctx.GetSender())

	return nil
}
