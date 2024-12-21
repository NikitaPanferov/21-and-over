package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) getGameState() *entities.GameState {
	players := make([]*entities.Player, 0, len(s.players))
	for _, player := range s.players {
		preparedCards := make([]*entities.Card, 0, len(player.Hand.Cards))
		for _, card := range s.dealer.Cards {
			preparedCard := *card
			if preparedCard.IsHidden {
				preparedCard.Rank = ""
				preparedCard.Suit = ""
			}

			preparedCards = append(preparedCards, &preparedCard)
		}

		player.Hand.Cards = preparedCards
		players = append(players, player)

	}

	return &entities.GameState{
		Players:        players,
		ActivePlayerIP: s.activePlayerIP,
		DillerHand:     s.dealer,
	}
}

func (s *Service) GetState(ctx context.Context) entities.State {
	return s.state
}

// Переход хода к следующему игроку
func (s *Service) switchToNextPlayer(currentPlayerIP string) bool {
	playerFound := false
	for _, player := range s.players {
		if playerFound {
			s.activePlayerIP = player.IP
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
	if s.activePlayerIP != playerIP {
		return nil, entities.ErrNotYourTurn
	}

	// Проверка наличия игрока
	player, ok := s.players[playerIP]
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
	return s.activePlayerIP
}

func (s *Service) Reset(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_ = s.playerRepo.SavePlayers(s.getGameState().Players)

	s.players = make(map[string]*entities.Player)
	s.deck = entities.NewDeck()
	s.activePlayerIP = ""
	s.dealer = &entities.Hand{}
	s.state = entities.JoinState
}
