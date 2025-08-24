package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/0x2A9/ai-resolver-mind-ms/data"
	"github.com/0x2A9/ai-resolver-mind-ms/services"
)

func SolveSudoku(c *fiber.Ctx) error {
	sudoku := new(data.Sudoku)

	if err := c.BodyParser(sudoku); err != nil {
		return c.JSON(err.Error())
	}

	services.Solve(&sudoku.Board, sudoku.SquareSize)
	return c.JSON(sudoku)
}
