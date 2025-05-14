package app

import (
	"context"
	"database/sql"
)

const (
	createTableVotesQuery = `
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

	createTableRatingsQuery = `
	CREATE TABLE IF NOT EXISTS ratings (
		image_name TEXT PRIMARY KEY,
		rating     INTEGER NOT NULL
	);`

	insertTableVotesQuery = `
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
	RETURNING id, created_at;`

	insertTableRatingsQuery = `
		INSERT INTO ratings(image_name, rating)
		VALUES ($1, $2)
        ON CONFLICT (image_name)
        DO UPDATE SET rating = EXCLUDED.rating;`
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) GetAllTableRatings(ctx context.Context) ([]RatingsModel, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT image_name, rating FROM ratings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []RatingsModel
	for rows.Next() {
		var m RatingsModel
		if err := rows.Scan(&m.Image, &m.Elo); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, nil
}

func (r *ProjectRepository) InsertTableVotes(ctx context.Context, votesModel *VotesModel) error {
	return r.DB.QueryRowContext(
		ctx,
		insertTableVotesQuery,
		votesModel.UserName,
		votesModel.ImageA,
		votesModel.ImageB,
		votesModel.ImageWinner,
		votesModel.ImageLoser,
		votesModel.EloWinnerPrevious,
		votesModel.EloWinnerNew,
		votesModel.EloLoserPrevious,
		votesModel.EloLoserNew,
	).Scan(&votesModel.ID, &votesModel.CreatedAt)
}

func (r *ProjectRepository) InsertTableRatings(ctx context.Context, ratings ...RatingsModel) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, rating := range ratings {
		if _, err := tx.ExecContext(
			ctx,
			insertTableRatingsQuery,
			rating.Image,
			rating.Elo,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func InitTableVotes(db *sql.DB) error {
	_, err := db.Exec(createTableVotesQuery)
	return err
}

func InitTableRatings(db *sql.DB) error {
	_, err := db.Exec(createTableRatingsQuery)
	return err
}
