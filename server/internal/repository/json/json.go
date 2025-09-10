package json

import (
	"encoding/json"
	"os"
	"sync"
)

type Repo struct {
	players      map[string]int
	mu           *sync.RWMutex
	initialChips int
	filepath     string
}

func New(initialChips int, filepath string) *Repo {
	repo := &Repo{
		players:      make(map[string]int),
		mu:           &sync.RWMutex{},
		initialChips: initialChips,
		filepath:     filepath,
	}

	if err := repo.loadFromFile(); err != nil {
		repo.players = make(map[string]int)
	}

	return repo
}

func (r *Repo) loadFromFile() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.Open(r.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&r.players); err != nil {
		return err
	}

	return nil
}
