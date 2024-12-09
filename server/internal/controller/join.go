package controller

import (
	"encoding/json"
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

	_ = entities.NewPlayer(request.Name)

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

	writeErr := ctx.Write(tcpserver.Response{
		Code: tcpserver.ResponseCodeSuccess,
	})
	if writeErr != nil {
		return fmt.Errorf("ctx.Write: %w", err)
	}

	//TODO: сделать норм логгер
	fmt.Printf("INFO: handled join request for: %s\n", ctx.GetSender())

	return nil
}
