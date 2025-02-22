package robotname

import "math/rand"

// Define the Robot type here.
type Robot struct {
	name             string
	unavailableNames *[]string
}

func randomName() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")
	b := make([]rune, 5)
	for i := range b {
		if i < 2 {
			b[i] = letters[rand.Intn(len(letters))]
		} else {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(b)
}

func (r *Robot) Name() (string, error) {
	return randomName(), nil
}

func (r *Robot) Reset() {
	panic("Please implement the Reset function")
}
