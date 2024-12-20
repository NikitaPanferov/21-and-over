package game

import (
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func (s *Service) Join(ctx tcpserver.Context, player entities.Player) error {
	if s.IsGameOn {
		return entities.ErrJoinGameIsOn
	}

	if len(s.Players) == MaxPlayers {
		return entities.ErrJoinGameIsFull
	}

	if _, ok := s.Players[ctx.GetSender()]; ok {
		return entities.ErrJoinPlayerAlreadyInGame
	}

	s.Players[ctx.GetSender()] = &player
	return nil
}
