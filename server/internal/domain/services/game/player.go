package game

import "github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"

func (s *Service) GetPlayer(
	name, ip string,
) *entities.Player {
	playerChips := s.playerRepo.GetPlayer(name)

	return &entities.Player{
		Name:  name,
		IP:    ip,
		Chips: playerChips,
		Hands: make([]*entities.Hand, 0),
	}
}
