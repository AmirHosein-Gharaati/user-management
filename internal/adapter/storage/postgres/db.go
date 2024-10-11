package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	DB  *sql.DB
	url string
}

func New(config *config.DB) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		slog.Error("could not connect to database", "error", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("could not ping to database1", "error", err)
		return nil, err
	}
	slog.Info("connected to postgres database")

	return &DB{
		DB:  db,
		url: url,
	}, nil
}

func (db *DB) Migrate() error {
	m, err := migrate.New("file://internal/adapter/storage/postgres/migrations", db.url)
	if err != nil {
		slog.Error("migration failed", "error", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("migration up failed", "error", err)
		return err
	}

	return nil
}
