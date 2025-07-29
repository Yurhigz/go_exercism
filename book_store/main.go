package bookstore

bookReduction := map[int]int{
	1 : 800,
	2 : 1520,
	3 : 2160,
	4 : 2560,
	5 : 3000,
}

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}
	booklist := map[int]int{}
	for _,book := range books {
		booklist[book] += 1
	}


    

    
	panic("Please implement the Cost function")
}
