package game

import "github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"

const (
	//TODO: убрать в конфиг
	MaxPlayers = 5
)

type Service struct {
	Players  map[string]*entities.Player
	Deck     *entities.Deck
	IsGameOn bool
}

func New() *Service {
	return &Service{
		Players: make(map[string]*entities.Player),
		Deck:    entities.NewDeck(),
		IsGameOn: false,
	}
}
