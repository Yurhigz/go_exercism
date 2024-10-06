package diffsquares

func SquareOfSum(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += i
	}
	return total * total
}

func SumOfSquares(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += i * i
	}
	return total
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
