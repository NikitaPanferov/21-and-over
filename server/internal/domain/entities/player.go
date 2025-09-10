package entities

type Player struct {
	Hand  *Hand  `json:"hand"`
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Chips int    `json:"chips"`
	Bet   int    `json:"bet"`
}

func NewPlayer(name string, ip string) *Player {
	return &Player{
		Hand: &Hand{},
		Name: name,
		IP:   ip,
	}
}
