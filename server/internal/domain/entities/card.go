package entities

type Suit string

const (
	Spades   Suit = "spades"
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Clubs    Suit = "clubs"
)

type Rank string

const (
	Ace   Rank = "ace"
	Two   Rank = "two"
	Three Rank = "three"
	Four  Rank = "four"
	Five  Rank = "five"
	Six   Rank = "six"
	Seven Rank = "seven"
	Eight Rank = "eight"
	Nine  Rank = "nine"
	Ten   Rank = "ten"
	Jack  Rank = "jack"
	Queen Rank = "queen"
	King  Rank = "king"
)

type Card struct {
	Suit     Suit
	Rank     Rank
	IsHidden bool
}
