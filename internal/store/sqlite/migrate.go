package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func MigrateUp(ctx context.Context, db *sql.DB) error {

	goose.SetBaseFS(migrationsFS)
	goose.SetDialect("sqlite3")

	if err := goose.UpContext(ctx, db, "migrations"); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
