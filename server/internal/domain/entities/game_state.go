package entities

type GameState struct {
	Players      []*Player `json:"players"`
	Deck         *Deck     `json:"deck"`
	ActivePlayer string    `json:"activePlayer"`
	DillerHand   *Hand     `json:"dillerHand"`
	State        bool      `json:"state"`
}
