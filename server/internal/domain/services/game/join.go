package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) Join(ctx context.Context, player *entities.Player) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.players) >= s.maxPlayers {
		return s.getGameState(), entities.ErrGameIsFull
	}

	if _, ok := s.players[player.IP]; ok {
		return s.getGameState(), entities.ErrPlayerAlreadyInGame
	}

	s.players[player.IP] = player

	if len(s.players) == s.maxPlayers {
		s.state = entities.WaitBetState
	}

	return s.getGameState(), nil
}
