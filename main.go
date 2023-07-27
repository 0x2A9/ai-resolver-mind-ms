package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/lamia-mortis/ai-resolver-mind-ms/controllers"
)

func main() {
    app := fiber.New()

	host, found := os.LookupEnv("APP_HOST")

    if !found {        
		host = "ai-resolver.mind-ms"
    }

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://ai-resolver.main, http://ai-resolver-main, https://ai-resolver.main, https://ai-resolver-main",
	}))
	
    setUpApiRoutes(app)
    app.Listen(host + ":8888")
}

func setUpApiRoutes(app *fiber.App) {
	app.Post("/api/sudoku/solve", controllers.SolveSudoku)
}
