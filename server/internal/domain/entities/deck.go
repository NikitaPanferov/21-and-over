package entities

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []*Card `json:"cards"`
}

func NewDeck() *Deck {
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	// Create a deck with two standard decks combined
	cards := make([]*Card, 0, 104) // 52 cards per deck * 2 decks
	for i := 0; i < 2; i++ {       // Add two decks
		for _, suit := range suits {
			for _, rank := range ranks {
				cards = append(cards, &Card{Suit: suit, Rank: rank, IsHidden: false})
			}
		}
	}

	// Shuffle the deck
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return &Deck{Cards: cards}
}
