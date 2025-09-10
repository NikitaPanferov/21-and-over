package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) GetPlayer(ctx context.Context, name, ip string) *entities.Player {
	playerChips := s.playerRepo.GetPlayer(name)

	return &entities.Player{
		Name:  name,
		IP:    ip,
		Chips: playerChips,
		Hand:  &entities.Hand{},
	}
}
