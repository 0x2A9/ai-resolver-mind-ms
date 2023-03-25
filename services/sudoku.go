package services 

import (
	"math"
)

func Solve(board *[9][9]int8, size int8) bool {
	return fillBoard(board, size)
}

func fillBoard(board *[9][9]int8, size int8) bool {
	xEmptyIndex, yEmptyIndex := findEmpty(board)

	if yEmptyIndex == -1 {
		return true
	}

	var numberToAdd int8

	for numberToAdd = 1; numberToAdd <= 9; numberToAdd++ {
		isValid := validate(board, size, numberToAdd, xEmptyIndex, yEmptyIndex)

		if isValid {
			board[yEmptyIndex][xEmptyIndex] = numberToAdd
			if fillBoard(board, size) {
				return true
			}

			board[yEmptyIndex][xEmptyIndex] = 0
		}
	}
	return false
}

func findEmpty(board *[9][9]int8) (int8, int8) {
	for yIndex := 0; yIndex < len(board); yIndex++ {
		row := board[yIndex]

		for xIndex := 0; xIndex < len(row); xIndex++ {
			if row[xIndex] == 0 {
				return int8(xIndex), int8(yIndex)
			}
		}
	}
	return -1, -1
}

func validate(board *[9][9]int8, size int8, numberToAdd int8, xEmptyIndex int8, yEmptyIndex int8) bool {
	isValid := true

	if contains(&board[yEmptyIndex], numberToAdd) {
		isValid = false
	}

	for _, row := range board {
		if row[xEmptyIndex] == numberToAdd {
			isValid = false
			break
		}
	}

	xEmptyIndexFl64 := float64(xEmptyIndex)
	yEmptyIndexFl64 := float64(yEmptyIndex)
	sizeFl64 := float64(size)

	xSquare := int8(math.Ceil((xEmptyIndexFl64 + 1) / sizeFl64))
	ySquare := int8(math.Ceil((yEmptyIndexFl64 + 1) / sizeFl64))

	skipXCells := (xSquare - 1) * size
	skipYCells := (ySquare - 1) * size

	initialX := skipXCells
	initialY := skipYCells

	for i := size; i > 0; i-- {
		for j := size; j > 0; j-- {
			if board[initialY][initialX] == numberToAdd && initialY != yEmptyIndex && initialX != xEmptyIndex {
				isValid = false
				break
			}
			initialY++
		}

		initialY = skipYCells
		initialX++
	}

	return isValid
}

func contains(haystack *[9]int8, needle int8) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
