package repository

import (
	"context"
	"fmt"

	"github.com/nishant007-tech/GoDig/internal/database"
	"github.com/nishant007-tech/GoDig/internal/logger"
	"go.uber.org/dig"
)

// User entity
type User struct {
	ID, Name string
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
}

// Parameters for NewUserRepository: Dig will fill fields by type and name.
type userRepoParams struct {
	dig.In

	DB     database.Connector `name:"primaryDB"`
	Logger *logger.Logger     `optional:"true"`
}

type userRepoImpl struct {
	db  database.Connector
	log *logger.Logger
}

// NewUserRepository returns an implementation wired with `primaryDB`.
func NewUserRepository(p userRepoParams) UserRepository {
	return &userRepoImpl{db: p.DB, log: p.Logger}
}

func (r *userRepoImpl) GetAll(ctx context.Context) ([]User, error) {
	if r.log != nil {
		r.log.Info("repository: connecting to DB")
	}
	if err := r.db.Connect(ctx); err != nil {
		return nil, fmt.Errorf("connect error: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM users")
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}
