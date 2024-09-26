package darts

func Score(x, y float64) int {
	s := x*x + y*y
	switch {
	case s <= 1:
		return 10
	case s > 1 && s <= 25:
		return 5
	case s > 25 && s <= 100:
		return 1
	default:
		return 0
	}
}
