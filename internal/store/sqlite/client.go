package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type Client struct {
	db *sql.DB
}

func (c *Client) migrate(ctx context.Context) error {
	migration, err := migrationsFS.ReadFile("migrations/001_init.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration: %w", err)
	}

	if _, err := c.db.ExecContext(ctx, string(migration)); err != nil {
		return fmt.Errorf("failed to exec migrations: %w", err)
	}

	return nil
}

func NewClient(ctx context.Context, dbPath string) (*Client, error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failded to ping database: %w", err)
	}
	client := &Client{db: db}

	if err := client.migrate(ctx); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	slog.InfoContext(ctx, "Sqlite connected", "path", dbPath)
	return client, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
func (c *Client) DB() *sql.DB {
	return c.db
}
