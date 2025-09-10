package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
	"github.com/NikitaPanferov/21-and-over/server/pkg/logger"
)

func (s *Service) handleEndGame(ctx context.Context) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.state = entities.ResultState

	for s.DoesDealerNeedCards() {
		card := s.deck.DrawTopCard()
		if card == nil {
			return s.getGameState(), entities.ErrDeckEmpty
		}
		s.dealer.Cards = append(s.dealer.Cards, card)
	}

	for _, card := range s.dealer.Cards {
		card.IsHidden = false
	}

	for _, player := range s.players {
		player.Chips += player.Bet
		player.Bet = 0
	}

	s.calculateWinnersAndPayouts(ctx)

	return s.getGameState(), nil
}

func (s *Service) DoesDealerNeedCards() bool {
	scores := s.dealer.CalculateScores()
	for _, score := range scores {
		if score < 17 {
			return true
		}
	}

	return false
}

func (s *Service) calculateWinnersAndPayouts(ctx context.Context) {
	dillerScores := s.dealer.CalculateScores()
	dillerBestScore := getBestScore(dillerScores)

	for _, player := range s.players {
		playerScores := player.Hand.CalculateScores()
		playerBestScore := getBestScore(playerScores)

		switch {
		case playerBestScore > 21:
			logger.InfoContext(ctx, "player lost", "playerName", player.Name, "bet", player.Bet)
			break
		case dillerBestScore > 21:
		case playerBestScore > dillerBestScore:
			handleWin(player)
			logger.InfoContext(ctx, "player win", "playerName", player.Name, "bet", player.Bet)
			break
		default:
			logger.InfoContext(ctx, "player lost", "playerName", player.Name, "bet", player.Bet)
			break
		}
	}
}

// Выбирает лучший результат (максимум <= 21), или минимальный результат, если все > 21
func getBestScore(scores []int) int {
	bestScore := 0
	for _, score := range scores {
		if score > bestScore && score <= 21 {
			bestScore = score
		}
	}
	if bestScore == 0 {
		// Если все результаты > 21, берем минимальный
		bestScore = scores[0]
		for _, score := range scores {
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

// Вычисляет выигрыш для игрока
func calculatePayout(player *entities.Player) int {
	// Пример: игрок получает 1.5x от ставки за выигрыш
	return int(float64(player.Bet) * 1.5)
}

func handleWin(player *entities.Player) {
	player.Chips += calculatePayout(player)
}
