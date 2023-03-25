package data

type Sudoku struct {
	Board        [9][9]int8  `json:"board"`
	SquareSize   int8        `json:"squareSize"`
}