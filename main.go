package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/lamia-mortis/ai-resolver-mind-ms/handlers"
)

func main() {
	app := fiber.New()

	host, found := os.LookupEnv("APP_HOST")

	if !found {
		host = "ai-resolver-mind-ms"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://ai-resolver.main, http://ai-resolver-main, https://ai-resolver.main, https://ai-resolver-main",
	}))

	setUpApiRoutes(app)
	err := app.Listen(host + ":8888")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func setUpApiRoutes(app *fiber.App) {
	app.Post("/api/sudoku/solve", handlers.SolveSudoku)
}
