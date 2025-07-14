package database

import (
	"context"
	"database/sql"
	"fmt"
)

type MySQLConnector struct {
	DSN string
	DB  *sql.DB
}

func (m *MySQLConnector) Connect(ctx context.Context) error {
	db, err := sql.Open("mysql", m.DSN)
	if err != nil {
		return fmt.Errorf("mysql open: %w", err)
	}
	m.DB = db
	return m.DB.PingContext(ctx)
}

func (m *MySQLConnector) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.QueryContext(ctx, query, args...)
}

func (m *MySQLConnector) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.DB.ExecContext(ctx, query, args...)
}

// NewMySQL returns a Connector under `name:"secondaryDB"`.
func NewMySQL() (Connector, error) {
	return &MySQLConnector{
		DSN: "user:pass@tcp(localhost:3306)/dbname",
	}, nil
}
