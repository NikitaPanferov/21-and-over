package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) getGameState() *entities.GameState {
	players := make([]*entities.Player, 0, len(s.Players))
	for _, player := range s.Players {
		players = append(players, player)
	}

	return &entities.GameState{
		Players:      players,
		Deck:         s.Deck,
		ActivePlayerIP: s.ActivePlayerIP,
		DillerHand:   s.DillerHand,
	}
}

func (s *Service) GetState(ctx context.Context) entities.State {
	return s.state
}

// Переход хода к следующему игроку
func (s *Service) switchToNextPlayer(currentPlayerIP string) bool {
	playerFound := false
	for _, player := range s.Players {
		if playerFound {
			s.ActivePlayerIP = player.IP
			return true
		}
		if player.IP == currentPlayerIP {
			playerFound = true
		}
	}

	return false
}

func (s *Service) validatePlayerTurn(playerIP string) (*entities.Player, error) {
	// Проверка активного игрока
	if s.ActivePlayerIP != playerIP {
		return nil, entities.ErrNotYourTurn
	}

	// Проверка наличия игрока
	player, ok := s.Players[playerIP]
	if !ok {
		return nil, entities.ErrPlayerNotFound
	}

	// Проверка завершенности хода
	if player.Hand.Done {
		return nil, entities.ErrPlayersHandIsDone
	}

	// Проверка текущих очков
	playerScores := player.Hand.CalculateScores()
	if !isBusted(playerScores) {
		player.Hand.Done = true
		s.switchToNextPlayer(playerIP)
		return nil, entities.ErrPlayerShouldBeAlreadyDone
	}

	return player, nil
}

// Проверка не проиграл ли игрок
func isBusted(scores []int) bool {
	for _, score := range scores {
		if score <= 21 {
			return false
		}
	}
	return true
}

func (s *Service) GetActivePlayerIP(ctx context.Context) string {
	return s.ActivePlayerIP
}
