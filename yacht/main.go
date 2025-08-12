package yacht

func numbersCount(value int, dice []int) int {
	var points int
	for _, v := range dice {
		if v == value {
			points += value
		}
	}
	return points
}

func countOccurences(dice []int) map[int]int {
	occurences := make(map[int]int)
	for _, v := range dice {
		occurences[v] += 1
	}
	return occurences
}

func sum(dice []int) int {
	total := 0
	for _, v := range dice {
		total += v
	}
	return total
}

func Score(dice []int, category string) int {
	points := 0
	occurrences := countOccurences(dice)

	switch category {
	case "ones":
		points = numbersCount(1, dice)
	case "twos":
		points = numbersCount(2, dice)
	case "threes":
		points = numbersCount(3, dice)
	case "fours":
		points = numbersCount(4, dice)
	case "fives":
		points = numbersCount(5, dice)
	case "sixes":
		points = numbersCount(6, dice)

	case "full house":
		if len(occurrences) == 2 {
			for i, v := range occurrences {
				if v == 3 || v == 2 {
					points += i * v
				}
			}
		}

	case "four of a kind":
		for i, v := range occurrences {
			if v >= 4 {
				points = i * 4
			}
		}

	case "little straight":
		if len(occurrences) == 5 && occurrences[1] == 1 && occurrences[2] == 1 && occurrences[3] == 1 && occurrences[4] == 1 && occurrences[5] == 1 {
			points = 30
		}

	case "big straight":
		if len(occurrences) == 5 && occurrences[2] == 1 && occurrences[3] == 1 && occurrences[4] == 1 && occurrences[5] == 1 && occurrences[6] == 1 {
			points = 30
		}

	case "choice":
		points = sum(dice)

	case "yacht":
		if len(occurrences) == 1 {
			points = 50
		}
	}
	return points
}
