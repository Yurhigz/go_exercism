package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func isMatrix(s string) bool {
	rows := strings.Split(s, "\n")
	if len(rows) == 0 {
		return false
	}
	var col int
	for i, values := range rows {
		nums := strings.Fields(values)
		if i == 0 {
			col = len(nums)
		} else if len(nums) != col {
			return false
		}
		for _, v := range nums {
			_, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func New(s string) (Matrix, error) {
	if !isMatrix(s) {
		return Matrix{}, errors.New("argument must be a matrix")
	}
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))

	for i, value := range rows {
		nums := strings.Fields(value)
		for _, v := range nums {
			numbers, _ := strconv.Atoi(v)
			matrix[i] = append(matrix[i], numbers)
		}
	}
	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for _, l := range m {
		for k, v := range l {
			cols[k] = append(cols[k], v)
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, l := range m {
		rows[i] = append(rows[i], l...)
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}

func main() {
	// test1 := "1 2\n10 20"
	// test2 := "89 1903 3\n18 3 1\n9 4 800"
	test3 := "1 2 3\n4 5 6\n7 8 9\n 8 7 6"
	m, _ := New(test3)

	m.Cols()
	m.Set(0, 0, 10)
	fmt.Println(m)
}
