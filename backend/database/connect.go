package database

import (
	"database/sql"
	"fmt"
	"web-imagecomparison/env"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *env.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to DB: %w", err)
	}

	return db, nil
}
