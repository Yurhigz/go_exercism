package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows    [][]int
	columns [][]int
}

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
	cols := [][]int{}
	rows := [][]int{}
	iteratorCol := 0
	iteratorRow := 0
	col := []int{}
	row := []int{}
	for _, value := range s {
		if value == '\n' {
			iteratorRow++
			cols = append(cols, col)
			rows = append(rows, row)
			col = []int{}
			row = []int{}
		} else {
			col = append(col, int(value))
			row = append(row, int(value))
			iteratorCol++
		}
	}
	return Matrix{rows: rows, columns: cols}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	return m.columns
}

func (m Matrix) Rows() [][]int {
	return m.rows
}

func (m Matrix) Set(row, col, val int) bool {
	m.columns[row][col] = val
	m.rows[col][row] = val
	return true
}

func main() {
	// test1 := "1 2\n10 20"
	test2 := "89 1903 3\n18 3 1\n9 4 800"
	// fmt.Println(New(test1))
	// cols := [][]int{}
	// rows := [][]int{}
	// iteratorCol := 0
	// iteratorRow := 0
	// col := []int{}
	// row := []int{}
	// for _, value := range strings.Split(test2, "\n") {
	// 	fmt.Println("len vaut : ", len(strings.Split(value, " ")))
	// 	for _, v := range strings.Split(value, " ") {
	// 		fmt.Println(v)
	// 	}
	// }
	fmt.Println(isMatrix(test2))
}
