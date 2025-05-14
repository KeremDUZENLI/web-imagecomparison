package app

import (
	"context"
	"database/sql"
	"web-imagecomparison/env"
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
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) InsertTableVotes(ctx context.Context, v *VoteModel) error {
	return r.DB.QueryRowContext(
		ctx,
		insertTableVotesQuery,
		v.UserName,
		v.ImageA,
		v.ImageB,
		v.ImageWinner,
		v.ImageLoser,
		v.EloWinnerPrevious,
		v.EloWinnerNew,
		v.EloLoserPrevious,
		v.EloLoserNew,
	).Scan(&v.ID, &v.CreatedAt)
}

func (r *ProjectRepository) GetAllTableRatings(ctx context.Context) (map[string]int, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT image_name, rating FROM ratings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[string]int)
	for rows.Next() {
		var img string
		var rt int
		if err := rows.Scan(&img, &rt); err != nil {
			return nil, err
		}
		m[img] = rt
	}
	return m, nil
}

func (r *ProjectRepository) UpdateTableRatings(ctx context.Context, winner string, loser string, delta int) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	upsert := `
		INSERT INTO ratings(image_name, rating)
		VALUES ($1, $2)
		ON CONFLICT (image_name) DO NOTHING;`
	if _, err := tx.ExecContext(ctx, upsert, winner, env.DEFAULT_RATING); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, upsert, loser, env.DEFAULT_RATING); err != nil {
		return err
	}

	update := `
		UPDATE ratings
		SET rating = CASE
			WHEN image_name = $1 THEN rating + $3
			WHEN image_name = $2 THEN rating - $3
		END
		WHERE image_name IN ($1, $2);`
	if _, err := tx.ExecContext(ctx, update, winner, loser, delta); err != nil {
		return err
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
