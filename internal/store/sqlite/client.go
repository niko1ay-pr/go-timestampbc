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

func NewClient(ctx context.Context, dbPath string, opts ...Option) (*Client, error) {
	var o options
	for _, fn := range opts {
		fn(&o)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failded to ping database: %w", err)
	}
	client := &Client{db: db}

	if !o.skipMigrations {
		if err := MigrateUp(ctx, db); err != nil {
			db.Close()
			return nil, fmt.Errorf("migration failed: %w", err)
		}
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
