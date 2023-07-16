package auth

import (
	"fmt"

	"github.com/ennwy/auth/internal/app"
	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
)

func (s *Service) SignUp(ctx *fiber.Ctx) error {
	user := app.User{}

	if err := ctx.BodyParser(&user); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	if err := user.HashPassword(); err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user.GenerateActivationLink()

	if err := s.Storage.CreateUser(ctx.Context(), user); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (s *Service) SignIn(ctx *fiber.Ctx) error {
	return nil
}

func (s *Service) Logout(ctx *fiber.Ctx) error {
	return nil
}

func (s *Service) Activate(ctx *fiber.Ctx) error {
	return nil
}

func (s *Service) Refresh(ctx *fiber.Ctx) error {
	return nil
}

func (s *Service) ListUsers(ctx *fiber.Ctx) error {
	users, err := s.Storage.ListUsers(ctx.Context())
	if err != nil {
		return fmt.Errorf("handler: listing: %w", err)
	}

	l.Info("List users:", users)

	if err = ctx.JSON(users); err != nil {
		return fmt.Errorf("list users: json: %w", err)
	}

	return nil
}
