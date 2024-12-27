package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Junx27/inventory-golang/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg config.Config) error {
	// Escape password to handle special characters
	pass := url.QueryEscape(cfg.DatabasePassword)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DatabaseUsername,
		pass,
		fmt.Sprintf("%s:%s", cfg.DatabaseHost, cfg.DatabasePort),
		cfg.DatabaseName,
	)

	// Open a database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	// Ensure migration path exists
	if _, err := os.Stat(cfg.MigrationPath); os.IsNotExist(err) {
		return fmt.Errorf("migration path does not exist: %s", cfg.MigrationPath)
	}

	// Create migration driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("Failed to create migration driver: %v", err)
		return err
	}

	// Initialize migration instance
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", cfg.MigrationPath),
		"postgres", driver)
	if err != nil {
		log.Printf("Failed to initialize migration: %v", err)
		return err
	}

	// Run migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
		} else {
			log.Printf("Migration failed: %v", err)
			return err
		}
	}

	log.Println("Migrations applied successfully.")
	return nil
}
