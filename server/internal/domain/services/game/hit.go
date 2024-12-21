package game

import (
	"context"
	"fmt"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) Hit(ctx context.Context, playerIP string) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Валидация хода игрока
	player, err := s.validatePlayerTurn(playerIP)
	if err != nil {
		return s.getGameState(), fmt.Errorf("s.validatePlayerTurn: %w", err)
	}

	// Взятие карты
	card := s.deck.DrawTopCard()
	if card == nil {
		return s.getGameState(), entities.ErrDeckEmpty
	}
	card.IsHidden = false
	player.Hand.Cards = append(player.Hand.Cards, card)

	// Проверка очков после взятия карты
	playerScores := player.Hand.CalculateScores()
	if !isBusted(playerScores) {
		player.Hand.Done = true
		switched := s.switchToNextPlayer(playerIP)
		if !switched {
			s.handleEndGame(ctx)
			return s.getGameState(), nil
		}
	}

	return s.getGameState(), nil
}
