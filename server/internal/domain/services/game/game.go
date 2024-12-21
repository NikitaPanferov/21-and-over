package game

import (
	"sync"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

type PlayerRepo interface {
	GetPlayer(name string) int
	SavePlayers(players []*entities.Player) error
}

type Service struct {
	players        map[string]*entities.Player
	deck           *entities.Deck
	activePlayerIP string
	dealer         *entities.Hand
	maxPlayers     int
	mu             *sync.RWMutex
	state          entities.State

	playerRepo PlayerRepo
}

func New(maxPlayers int, playerRepo PlayerRepo) *Service {
	return &Service{
		players:        make(map[string]*entities.Player),
		deck:           entities.NewDeck(),
		activePlayerIP: "",
		dealer:         &entities.Hand{},
		maxPlayers:     maxPlayers,
		mu:             &sync.RWMutex{},
		playerRepo:     playerRepo,
		state:          entities.JoinState,
	}
}
