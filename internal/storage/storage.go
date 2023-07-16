package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/ennwy/auth/internal/app"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn   *pgx.Conn
	config *DBConfig
}

var _ app.Storage = (*DB)(nil)

var l app.Logger = nil

func New(log app.Logger) *DB {
	l = log

	return &DB{
		config: NewDBConf(),
	}
}

func (db *DB) Connect(ctx context.Context) (err error) {
	if db.conn, err = pgx.Connect(ctx, db.config.getConnectString()); err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	return nil
}

func (db *DB) Close(ctx context.Context) error {
	if err := db.conn.Close(ctx); err != nil {
		return fmt.Errorf("storage close: %w", err)
	}

	return nil
}

type DBConfig struct {
	port     string
	host     string
	name     string
	user     string
	password string
}

func (db *DBConfig) getConnectString() string {
	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.host,
		db.port,
		db.user,
		db.password,
		db.name,
	)

	l.Debug("Database config: ", info)

	return info
}

func NewDBConf() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
	}
}
