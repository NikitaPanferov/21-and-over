package entities

type GameState struct {
	Players        []*Player `json:"players"`
	Deck           *Deck     `json:"deck"`
	ActivePlayerIP string    `json:"activePlayerIP"`
	DillerHand     *Hand     `json:"dillerHand"`
	State          State     `json:"state"`
}
