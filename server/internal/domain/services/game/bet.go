package game

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func (s *Service) Bet(ctx context.Context, playerIP string, bet int) (*entities.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	player, ok := s.players[playerIP]
	if !ok {
		return s.getGameState(), entities.ErrPlayerNotFound
	}

	if player.Chips < bet {
		return s.getGameState(), entities.ErrNotEnoughChips
	}

	if player.Bet != 0 {
		return s.getGameState(), entities.ErrPlayerHasAlreadyDidBet
	}

	player.Bet = bet
	player.Chips -= bet

	if s.checkIfAllBeted() {
		s.state = entities.PlayState
		s.giveInitialCards()
	}

	for _, player := range s.players {
		s.activePlayerIP = player.IP
		break
	}

	return s.getGameState(), nil
}

func (s *Service) checkIfAllBeted() bool {
	for _, player := range s.players {
		if player.Bet == 0 {
			return false
		}
	}

	return true
}

func (s *Service) giveInitialCards() {
	for _, player := range s.players {
		card := s.deck.DrawTopCard()
		card.IsHidden = false
		player.Hand.Cards = append(player.Hand.Cards, card)

		card = s.deck.DrawTopCard()
		card.IsHidden = false
		player.Hand.Cards = append(player.Hand.Cards, card)
	}

	card := s.deck.DrawTopCard()
	card.IsHidden = false
	s.dealer.Cards = append(s.dealer.Cards, s.deck.DrawTopCard())
}
