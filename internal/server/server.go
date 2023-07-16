package server

import (
	"context"
	"fmt"
	"net"

	"github.com/ennwy/auth/internal/app"
	"github.com/ennwy/auth/internal/auth"
	"github.com/ennwy/auth/internal/storage"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	ctx     context.Context
	App     *fiber.App
	Auth    *auth.Service
	Address string
}

var _ app.Server = (*Server)(nil)
var l app.Logger

func New(ctx context.Context, log app.Logger, c Config) (*Server, error) {
	l = log

	s := &Server{
		ctx: ctx,
		App: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		Address: net.JoinHostPort(c.Host, c.Port),
	}

	auth, err := auth.NewService(ctx, log, storage.New(log))
	if err != nil {
		return nil, fmt.Errorf("new service: %w", err)
	}
	s.Auth = auth

	l.Debug(s.Address)
	s.setRouters()

	return s, nil
}

func (s *Server) Start() error {
	return s.App.Listen(s.Address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.App.Server().ShutdownWithContext(ctx)
}

func (s *Server) Ctx() context.Context {
	return s.ctx
}

func (s *Server) setRouters() {
	s.App.Post("/signup", s.Auth.SignUp)
	s.App.Post("/signin", s.Auth.SignIn)
	s.App.Post("/logout", s.Auth.Logout)
	s.App.Get("/activate/:link", s.Auth.Activate)
	s.App.Get("/users", s.Auth.ListUsers)
}
