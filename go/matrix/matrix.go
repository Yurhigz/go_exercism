package matrix

// Define the Matrix type here.
type Matrix struct {
	rows    [][]int
	columns [][]int
}

func New(s string) (Matrix, error) {
	col := [][]int{}
	row := [][]int{}
	iteratorCol := 0
	iteratorRow := 0
	for value := range s {
		if value == '\n' {
			iteratorRow++
		} else {
			col[iteratorCol] = append(col[iteratorCol], int(value))
			row[iteratorRow] = append(row[iteratorRow], int(value))
			iteratorCol++
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
