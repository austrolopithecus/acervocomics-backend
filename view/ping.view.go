package view

import "github.com/gofiber/fiber/v3"

func Ping(c fiber.Ctx) error {
	return c.SendString("pong")
}
