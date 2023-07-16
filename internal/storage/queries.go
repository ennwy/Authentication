package storage

import (
	"context"
	"fmt"

	"github.com/ennwy/auth/internal/app"
)

const (
	qCreate        = `INSERT INTO person(email,password,activation_link) VALUES($1,$2,$3);`
	qCheckIfExists = `SELECT COUNT(*) FROM person WHERE email = $1;`
	qListUsers     = `SELECT id, email, password FROM person;`
)

func (db *DB) CreateUser(ctx context.Context, user app.User) error {
	_, err := db.conn.Exec(
		ctx,
		qCreate,

		user.Email,
		user.Password,
		user.ActivationLink,
	)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (db *DB) ListUsers(ctx context.Context) ([]app.User, error) {
	rows, err := db.conn.Query(ctx, qListUsers)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	users := make([]app.User, 0, 1)
	var user app.User

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		); err != nil {
			l.Error("list users: rows scan: %w", err)
			continue
		}

		users = append(users, user)
		user = app.User{}
	}

	return users, nil
}
