package game

import (
	"sync"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

type PlayerRepo interface {
	GetPlayer(name string) int
}

type Service struct {
	Players      map[string]*entities.Player
	Deck         *entities.Deck
	ActivePlayerIP string
	DillerHand   *entities.Hand
	MaxPlayers   int
	mu           *sync.RWMutex
	state        entities.State

	playerRepo PlayerRepo
}

func New(maxPlayers int, playerRepo PlayerRepo) *Service {
	return &Service{
		Players:      make(map[string]*entities.Player),
		Deck:         entities.NewDeck(),
		ActivePlayerIP: "",
		DillerHand:   &entities.Hand{},
		MaxPlayers:   maxPlayers,
		mu:           &sync.RWMutex{},
		playerRepo:   playerRepo,
		state:        entities.JoinState,
	}
}
