package controller

import (
	"encoding/json"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/controller/types"
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func (c *Controller) joinHandler(ctx *tcpserver.Context) error {
	fmt.Printf("Join handler received: %s\n", ctx.GetMessage().Data.(map[string]interface{}))

	var request types.JoinRequest
	err := json.Unmarshal(ctx.GetRawData(), &request)
	if err != nil {
		writeErr := ctx.ReplyWithError(tcpserver.CodeClientError, err)
		if writeErr != nil {
			return fmt.Errorf("ctx.ReplyWithError: %w", err)
		}

		//TODO: сделать норм логгер
		fmt.Printf("INFO: handled error: %v\n", err)

		return nil
	}

	_ = entities.NewPlayer(request.Name)

	err = ctx.Reply(tcpserver.CodeSuccess, nil)
	if err != nil {
		return fmt.Errorf("ctx.Reply: %w", err)
	}

	//TODO: сделать норм логгер
	fmt.Printf("INFO: handled join request for: %s (%s)\n", request.Name, ctx.GetSender())

	return nil
}
