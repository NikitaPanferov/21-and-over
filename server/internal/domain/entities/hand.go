package entities

type Hand struct {
	Cards []*Card `json:"cards"`
	Done  bool    `json:"done"`
}

func (h *Hand) CalculateScores() []int {
	total := 0
	aceCount := 0

	// Подсчет очков
	for _, card := range h.Cards {
		switch card.Rank {
		case Ace:
			aceCount++
			total += 1 // Считаем тузы как 1
		case Two:
			total += 2
		case Three:
			total += 3
		case Four:
			total += 4
		case Five:
			total += 5
		case Six:
			total += 6
		case Seven:
			total += 7
		case Eight:
			total += 8
		case Nine:
			total += 9
		case Ten, Jack, Queen, King:
			total += 10
		}
	}

	// Генерация всех возможных значений
	scores := []int{total}
	for i := 0; i < aceCount; i++ {
		newScores := []int{}
		for _, score := range scores {
			newScores = append(newScores, score+10) // Учитываем туз как 11
		}
		scores = append(scores, newScores...)
	}

	// Удаляем дубли и сортируем
	uniqueScores := map[int]struct{}{}
	for _, score := range scores {
		uniqueScores[score] = struct{}{}
	}

	validScores := []int{}
	for score := range uniqueScores {
		if score <= 21 {
			validScores = append(validScores, score)
		}
	}

	// Если есть значения <= 21, возвращаем их
	if len(validScores) > 0 {
		return validScores
	}

	// Если все значения больше 21, возвращаем все уникальные значения
	result := []int{}
	for score := range uniqueScores {
		result = append(result, score)
	}
	return result
}
