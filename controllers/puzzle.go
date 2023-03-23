package controllers

import "github.com/gofiber/fiber/v2"

func Solve(c *fiber.Ctx) error {
    return c.Send(c.Body())
}