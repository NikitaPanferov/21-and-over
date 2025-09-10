package entities_test

import (
	"testing"

	"github.com/NikitaPanferov/21-and-over/server/internal/domain/entities"
)

func TestCalculateScores(t *testing.T) {
	tests := []struct {
		name     string
		cards    []*entities.Card
		expected []int
	}{
		{
			name: "Hand without Aces",
			cards: []*entities.Card{
				{Rank: entities.Five, Suit: entities.Hearts},
				{Rank: entities.Seven, Suit: entities.Spades},
			},
			expected: []int{12},
		},
		{
			name: "Hand with one Ace",
			cards: []*entities.Card{
				{Rank: entities.Ace, Suit: entities.Hearts},
				{Rank: entities.Seven, Suit: entities.Spades},
			},
			expected: []int{8, 18},
		},
		{
			name: "Hand with two Aces",
			cards: []*entities.Card{
				{Rank: entities.Ace, Suit: entities.Hearts},
				{Rank: entities.Ace, Suit: entities.Spades},
				{Rank: entities.Nine, Suit: entities.Diamonds},
			},
			expected: []int{11, 21},
		},
		{
			name: "Hand with values exceeding 21",
			cards: []*entities.Card{
				{Rank: entities.King, Suit: entities.Hearts},
				{Rank: entities.Queen, Suit: entities.Spades},
				{Rank: entities.Jack, Suit: entities.Clubs},
			},
			expected: []int{30},
		},
		{
			name:     "Empty hand",
			cards:    []*entities.Card{},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := &entities.Hand{Cards: tt.cards}
			result := hand.CalculateScores()

			// Проверяем, что результат содержит все ожидаемые значения
			match := func(slice1, slice2 []int) bool {
				if len(slice1) != len(slice2) {
					return false
				}
				m := make(map[int]bool)
				for _, v := range slice1 {
					m[v] = true
				}
				for _, v := range slice2 {
					if !m[v] {
						return false
					}
				}
				return true
			}

			if !match(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
