package entities

type GameState struct {
	Players        []*Player `json:"players"`
	ActivePlayerIP string    `json:"activePlayerIP"`
	DillerHand     *Hand     `json:"dillerHand"`
	State          State     `json:"state"`
}
