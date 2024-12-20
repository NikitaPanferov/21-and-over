package entities

type Hand struct {
	Cards []*Card `json:"cards"`
	Done  bool    `json:"done"`
}
