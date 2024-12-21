package game

import (
	"context"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) Stand(ctx context.Context, playerIP string) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Валидация хода игрока
	player, err := s.validatePlayerTurn(playerIP)
	if err != nil {
		return s.getGameState(), fmt.Errorf("s.validatePlayerTurn: %w", err)
	}

	// Переход хода следующему игроку
	player.Hand.Done = true
	switched := s.switchToNextPlayer(playerIP)
	if !switched {
		s.state = entities.ResultState
		return s.getGameState(), nil
	}

	return s.getGameState(), nil
}
