package matrix

import (
	"errors"
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
	rows := strings.Split(s, "\n")
	row := make([][]int, len(rows))
	col := make([][]int, len(strings.Fields(rows[0])))
	for i, value := range rows {
		nums := strings.Fields(value)
		for k, v := range nums {
			numbers, _ := strconv.Atoi(v)
			row[i] = append(row[i], numbers)
			col[k] = append(col[k], numbers)
		}
	}
	return Matrix{rows: row, columns: col}, nil
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
