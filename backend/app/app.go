package app

import (
	"database/sql"
	"net/http"
)

type App struct {
	DB      *sql.DB
	Service *ProjectService
}

func RunApp(db *sql.DB) (*App, error) {
	if err := InitTable(db); err != nil {
		return nil, err
	}
	repo := NewProjectRepository(db)
	svc := NewProjectService(repo)
	return &App{DB: db, Service: svc}, nil
}

func (a *App) Routes() {
	ctrl := NewProjectController(a.Service)
	http.Handle("/", http.FileServer(http.Dir("../docs")))
	http.HandleFunc("/api/votes", ctrl.HandleEntry)
}
