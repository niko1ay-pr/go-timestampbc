package sqlite

import (
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
)

var migrationsFS embed.FS

type Client struct {
	db *sql.DB
}
