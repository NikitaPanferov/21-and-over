package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) Join(ctx context.Context, player *entities.Player) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.Players) >= s.MaxPlayers {
		return s.getGameState(), entities.ErrGameIsFull
	}

	if _, ok := s.Players[player.IP]; ok {
		return s.getGameState(), entities.ErrPlayerAlreadyInGame
	}

	s.Players[player.IP] = player

	if len(s.Players) == s.MaxPlayers {
		s.state = entities.WaitBetState
	}

	return s.getGameState(), nil
}