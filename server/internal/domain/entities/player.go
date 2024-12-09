package entities

type Player struct {
	Hands []*Hand
	Name  string
}

func NewPlayer(name string) *Player {
	return &Player{
		Hands: make([]*Hand, 0),
		Name:  name,
	}
}
