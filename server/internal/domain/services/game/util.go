package game

import "github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"

func (s *Service) getGameState() *entities.GameState {
	players := make([]*entities.Player, 0, len(s.Players))
	for _, player := range s.Players {
		players = append(players, player)
	}

	return &entities.GameState{
		Players:      players,
		Deck:         s.Deck,
		ActivePlayer: s.ActivePlayer,
		DillerHand:   s.DillerHand,
		
	}
}

func (s *Service) GetState() entities.State {
	return s.state
}
