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

func Score(dice []int, category string) int {
	var points int
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
		value := countOccurences(dice)
		if len(value) == 2 {
			for i, v := range value {
				if v == 3 || v == 2 {
					points += i * v
				}
			}
		} else {
			points = 0
		}

	case "four of a kind":
		value := countOccurences(dice)
		if len(value) <= 2 {
			for i, v := range value {
				if v >= 4 {
					points += 4 * i
				}
			}
		} else {
			points = 0
		}

	case "little straight":
		value := 0
		for _, v := range dice {
			value += v
		}
		if value == 15 {
			points = 30
		} else {
			points = 0
		}

	case "big straight":
		value := 0
		for _, v := range dice {
			value += v
		}
		if value == 20 {
			points = 30
		} else {
			points = 0
		}

	case "choice":
		for _, v := range dice {
			points += v
		}
	case "yacht":
		occurences := countOccurences(dice)
		if len(occurences) == 1 {
			points = 50
		} else {
			points = 0
		}
	}
	return points
}
