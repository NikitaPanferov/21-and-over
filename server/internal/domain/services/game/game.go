package game

import "github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"

type Service struct {
	Players map[string]*entities.Player
	Deck    *entities.Deck
}

func New() *Service {
	return &Service{
		Players: make(map[string]*entities.Player),
		Deck:    entities.NewDeck(),
	}
}
