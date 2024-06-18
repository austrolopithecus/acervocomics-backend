package view

import (
	"acervocomics/controller"
	"acervocomics/models"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

// Funções comuns de todos usuarios

// Rota de registro
func register(c fiber.Ctx) error {
	var body models.RegisterUserRequest

	if err := c.Bind().Body(&body); err != nil {
		res := models.CommonResponse{
			Message: fmt.Errorf("error binding request: %w", err).Error(),
			Succes:  false,
		}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}
	// Compare this snippet from controller/user.controller.go:
	registerResponse, err := controller.Register(&body)
	// Compare this snippet from controller/user.controller.go:
	if err != nil {
		res := models.CommonResponse{
			Message: fmt.Errorf("error registering user: %w", err).Error(),
			Succes:  false,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	// Compare this snippet from controller/user.controller.go:
	if registerResponse == nil {
		res := models.CommonResponse{
			Message: "User already exists",
			Succes:  false,
		}
		return c.Status(fiber.StatusConflict).JSON(res)
	}
	return c.Status(fiber.StatusCreated).JSON(registerResponse)
}

func login(c fiber.Ctx) error {
	return nil // TODO: Implement
}

func me(c fiber.Ctx) error {
	return nil // TODO: Implement
}

// Funções de administrador
// TODO: Adicionar middleware de admin
func listUsers(c fiber.Ctx) error {
	return nil
}
