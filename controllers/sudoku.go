package controllers

import (
    "github.com/gofiber/fiber/v2"

    "github.com/lamia-mortis/ai-resolver-mind-ms/services"
    "github.com/lamia-mortis/ai-resolver-mind-ms/data"
)

func SolveSudoku(c *fiber.Ctx) error {
    sudoku := new (data.Sudoku)

    if err := c.BodyParser(sudoku); err != nil {
        return err
    }

    services.Solve(&sudoku.Board, sudoku.SquareSize)
    return c.JSON(sudoku)
}