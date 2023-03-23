package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamia-mortis/ai-resolver-mind-ms/controllers"
)

func main() {
    app := fiber.New()

    setUpApiRoutes(app)
    app.Listen(":8888")
}

func setUpApiRoutes(app *fiber.App) {
	app.Post("/api/solve", controllers.Solve)
}
