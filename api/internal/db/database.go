package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/emilieanthony/senzr/internal/config"
	"github.com/jmoiron/sqlx"
)

const (
	driver         = "postgres"
	migrationsPath = "file://internal/db/migrations"
)

func newDBConnection() (_ *sqlx.DB, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("connection error: %v", err)
		}
	}()
	env, err := config.Env()
	if err != nil {
		return nil, fmt.Errorf("getting environment: %w", err)
	}
	db, err := sqlx.Connect(
		driver,
		fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
			env.DBUser, env.DBPassword, env.DBHost, env.DBName,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("connecting: %w", err)
	}
	return db, nil
}

func MigrateSchemas() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("migrating database: %w", err)
		}
	}()
	db, err := newDBConnection()
	if err != nil {
		return fmt.Errorf("connection: %w", err)
	}
	d, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		d,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		fmt.Println("No database changes")
		return nil
	}
	fmt.Println("Successfully migrated schemas")
	return nil
}
