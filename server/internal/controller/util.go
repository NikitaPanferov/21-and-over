package controller

import (
	"slices"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

var (
	stateActionsMapping = map[entities.State][]tcpserver.Action{
		entities.JoinState:    {tcpserver.ActionJoin},
		entities.WaitBetState: {tcpserver.ActionBet},
		entities.PlayState:    {tcpserver.ActionHit, tcpserver.ActionStand},
		entities.ResultState:  {},
	}
)

func (c *Controller) ValidateState(ctx *tcpserver.Context) bool {
	stateActions, ok := stateActionsMapping[c.gameService.GetState()]
	if !ok || !slices.Contains(stateActions, ctx.GetMessage().Action) {
		return false
	}

	return true
}
