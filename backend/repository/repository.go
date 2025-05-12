package repository

import (
	"database/sql"

	"web-imagecomparison/model"
)

const (
	createTableQuery = `
		CREATE TABLE IF NOT EXISTS votes (
			id SERIAL PRIMARY KEY,
			user_name TEXT,
			image_a TEXT,
			image_b TEXT,
			image_winner TEXT,
			image_loser TEXT,
			elo_winner_previous INTEGER,
			elo_winner_new INTEGER,
			elo_loser_previous INTEGER,
			elo_loser_new INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	insertRowQuery = `
		INSERT INTO votes (
			user_name, 
			image_a, 
			image_b, 
			image_winner, 
			image_loser,
			elo_winner_previous, 
			elo_winner_new, 
			elo_loser_previous, 
			elo_loser_new
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id, created_at`
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Insert(v *model.ProjectModel) error {
	return r.DB.QueryRow(
		insertRowQuery,
		v.UserName,
		v.ImageA,
		v.ImageB,
		v.ImageWinner,
		v.ImageLoser,
		v.EloWinnerPrevious,
		v.EloWinnerNew,
		v.EloLoserPrevious,
		v.EloLoserNew).
		Scan(&v.ID, &v.CreatedAt)
}

func InitTable(db *sql.DB) error {
	_, err := db.Exec(createTableQuery)
	return err
}
