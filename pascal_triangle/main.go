package main

import "fmt"

func Triangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	triangle := [][]int{}
	for i := 1; i <= n; i++ {

		if i == 1 {
			triangle = append(triangle, []int{1})
		} else {
			precLine := triangle[i-2]
			line := make([]int, i)
			line[0] = 1
			line[i-1] = 1
			for j := 1; j < i-1; j++ {
				line[j] = precLine[j-1] + precLine[j]
			}
			triangle = append(triangle, line)
		}

	}
	return triangle
}

func main() {
	test := []int{2}

	for _, v := range test {
		fmt.Println(Triangle(v))
	}
}

//     1
//    1 1
//   1 2 1
//  1 3 3 1
// 1 4 6 4 1
