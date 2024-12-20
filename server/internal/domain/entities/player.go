package entities

type Player struct {
	Hands []*Hand `hands:"json"`
	Name  string  `json:"name"`
	IP    string  `json:"-"`
	Chips int     `json:"chips"`
}

func NewPlayer(name string, ip string) *Player {
	return &Player{
		Hands: make([]*Hand, 0),
		Name:  name,
		IP:    ip,
	}
}
