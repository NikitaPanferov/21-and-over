package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (r *Repo) SavePlayers(players []*entities.Player) error {
	for _, player := range players {
		r.players[player.Name] = player.Chips
	}

	err := r.saveToFile()
	if err != nil {
		return fmt.Errorf("failed to save players to file: %w", err)
	}

	return nil
}

func (r *Repo) saveToFile() error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	file, err := os.OpenFile(r.filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(r.players); err != nil {
		return err
	}

	return nil
}
