package view

import "github.com/gofiber/fiber/v3"

func InitRoutes(c *fiber.App) {
	// Rotas de teste
	c.Get("/ping", Ping)
	// Rotas de usuário
	userGroup := c.Group("/user")
	userRoutes(userGroup)
}

func userRoutes(c fiber.Router) {
	// Funções comuns de todos usuarios
	c.Post("/register", register)

	c.Post("/login", login)
	c.Get("/me", me)
	// Rotas de admin
	c.Get("/list", listUsers)
}
