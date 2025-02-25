package matrix

import (
	"errors"
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
	cols := make([][]int, len(m))
	for _, l := range m {
		for k, v := range l {
			cols[k] = append(cols[k], v)
		}

	}
	return cols
}

func (m Matrix) Rows() [][]int {
	return m
}

func (m Matrix) Set(row, col, val int) bool {
	m.Rows()[row][col] = val
	return true
}
