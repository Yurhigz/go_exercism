package bookstore

func countBook(books []int) []int {
	booklist := make([]int, 5)
	for _, book := range books {
		if book >= 1 && book <= 5 {
			booklist[book-1] += 1
		}
	}
	return booklist
}

func generateCombinations(indices []int, size int) [][]int {
	if size == 0 {
		return [][]int{{}}
	}
	if len(indices) < size {
		return nil
	}

	var result [][]int
	for i := 0; i <= len(indices)-size; i++ {
		head := indices[i]
		rest := generateCombinations(indices[i+1:], size-1)
		for _, combo := range rest {
			result = append(result, append([]int{head}, combo...))
		}
	}
	return result
}

func possibleGroup(bookStock []int) [][]int {
	var result [][]int
	var indices []int

	// Étape 1 : récupérer les indices avec stock > 0
	for i, count := range bookStock {
		if count > 0 {
			indices = append(indices, i)
		}
	}

	maxSize := 5
	if len(indices) < maxSize {
		maxSize = len(indices)
	}
	for size := 1; size <= maxSize; size++ {
		comb := generateCombinations(indices, size)
		result = append(result, comb...)
	}
	return result

}

func Cost(books []int) int {
	bookReduction := map[int]int{
		1: 800,
		2: 1520,
		3: 2160,
		4: 2560,
		5: 3000,
	}
	if len(books) <= 0 {
		return 0
	}
	countBook := countBook(books)
	totalCost := 0
	selectedBook := []int{}
	for i := 1; i < 6; i++ {
		selectedBook = append(selectedBook, i)
	}

}
