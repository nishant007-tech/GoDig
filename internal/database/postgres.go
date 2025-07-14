package database

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresConnector struct {
	DSN string
	DB  *sql.DB
}

func (p *PostgresConnector) Connect(ctx context.Context) error {
	db, err := sql.Open("postgres", p.DSN)
	if err != nil {
		return fmt.Errorf("postgres open: %w", err)
	}
	p.DB = db
	return p.DB.PingContext(ctx)
}

// QueryContext forwards to the underlying *sql.DB.
func (p *PostgresConnector) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.DB.QueryContext(ctx, query, args...)
}

// ExecContext forwards to the underlying *sql.DB.
func (p *PostgresConnector) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.DB.ExecContext(ctx, query, args...)
}

// NewPostgres returns a Connector under `name:"primaryDB"`.
func NewPostgres() (Connector, error) {
	return &PostgresConnector{
		DSN: "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
	}, nil
}
