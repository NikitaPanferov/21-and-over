package sqlite

import "sync"

type Repo struct {
	players      map[string]int
	mu           *sync.RWMutex
	initialChips int
}

func New(initialChips int) *Repo {
	return &Repo{
		players:      make(map[string]int),
		initialChips: initialChips,
	}
}
