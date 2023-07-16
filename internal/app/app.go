package app

import (
	"context"
)

type Server interface {
	Start() error
	Stop(context.Context) error
}

type Logger interface {
	Debug(...any)
	Info(...any)
	Warn(...any)
	Error(...any)
	Fatal(...any)
}

type ConnectCloser interface {
	Connect(context.Context) error
	Close(context.Context) error
}

type Storage interface {
	ConnectCloser
	CreateUser(context.Context, User) error
	ListUsers(context.Context) ([]User, error)
}
