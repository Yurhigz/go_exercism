package robotname

import (
	"errors"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var robotNames = make(map[string]struct{}, 26*26*1000)

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
	if len(robotNames) >= 26*26*1000 {
		return "", errors.New("no name left")
	}
	if r.name == "" {
		test := randomName()
		if _, ok := robotNames[test]; ok {
			return r.Name()
		}
		r.name = test
		robotNames[test] = struct{}{}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

// // Package robotname is a small library for generating robot names.
// package robotname

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Robot is a struct with a string field.
// type Robot struct {
// 	name string
// }

// const maxRobotNames = 26 * 26 * 10 * 10 * 10

// var (
// 	namePool = generateRobotNames()
// 	idx      = 0
// )

// func generateRobotNames() []string {
// 	pos := 0
// 	names := make([]string, maxRobotNames)

// 	for i := 'A'; i <= 'Z'; i++ {
// 		for j := 'A'; j <= 'Z'; j++ {
// 			for k := 0; k < 1000; k++ {
// 				names[pos] = fmt.Sprintf("%s%s%03d", string(i), string(j), k)
// 				pos++
// 			}
// 		}
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
// 	return names
// }

// // Name returns the existing name or returns a newly assigned name.
// func (r *Robot) Name() (string, error) {
// 	if r.name != "" {
// 		return r.name, nil
// 	}

// 	if idx >= maxRobotNames {
// 		return "", fmt.Errorf("uniqueness exhausted")
// 	}

// 	r.name = namePool[idx]
// 	idx++

// 	return r.name, nil
// }

// // Reset erases the existing name without returning it to unused names.
// func (r *Robot) Reset() {
// 	r.name = ""
// }
