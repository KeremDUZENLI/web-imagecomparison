package app

import (
	"database/sql"
)

type App struct {
	DB      *sql.DB
	Service *ProjectService
}

func InitApp(db *sql.DB) (*App, error) {
	if err := InitTableVotes(db); err != nil {
		return nil, err
	}
	if err := InitTableRatings(db); err != nil {
		return nil, err
	}

	repo := NewProjectRepository(db)
	svc := NewProjectService(repo)
	return &App{DB: db, Service: svc}, nil
}
