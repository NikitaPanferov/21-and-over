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

	cards := make([]*Card, 0, 104)
	for i := 0; i < 2; i++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				cards = append(cards, &Card{Suit: suit, Rank: rank, IsHidden: true})
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return &Deck{Cards: cards}
}

func (d *Deck) DrawTopCard() *Card {
	if len(d.Cards) == 0 {
		return nil
	}

	topCard := d.Cards[0]

	d.Cards = d.Cards[1:]

	return topCard
}
