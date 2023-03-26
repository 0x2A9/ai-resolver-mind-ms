package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/lamia-mortis/ai-resolver-mind-ms/controllers"
)

func main() {
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://ai-resolver.main",
	}))
	
    setUpApiRoutes(app)
    app.Listen("ai-resolver.mind-ms:8888")
}

func setUpApiRoutes(app *fiber.App) {
	app.Post("/api/sudoku/solve", controllers.SolveSudoku)
}
