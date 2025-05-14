package app

import (
	"context"
	"database/sql"
)

const (
	createTableVotesQuery = `
		CREATE TABLE IF NOT EXISTS votes (
			user_name 			 TEXT,
			image_winner 		 TEXT,
			image_loser 		 TEXT,
			elo_winner_previous  INTEGER,
			elo_winner_new 		 INTEGER,
			elo_loser_previous 	 INTEGER,
			elo_loser_new 		 INTEGER,
			created_at 			 TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	createTableRatingsQuery = `
		CREATE TABLE IF NOT EXISTS ratings (
			image 	TEXT PRIMARY KEY,
			elo     INTEGER NOT NULL
		);`

	insertTableVotesQuery = `
		INSERT INTO votes (
			user_name,
			image_winner,
			image_loser,
			elo_winner_previous,
			elo_winner_new,
			elo_loser_previous,
			elo_loser_new
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING created_at;`

	insertTableRatingsQuery = `
		INSERT INTO ratings(image, elo)
		VALUES ($1, $2)
        ON CONFLICT (image)
        DO UPDATE SET elo = EXCLUDED.elo;`

	getDistinctUsersQuery = `
		SELECT DISTINCT user_name FROM votes;`

	getTableRatingsQuery = `
		SELECT image, elo FROM ratings;`
)

type ProjectRepository interface {
	GetAllUserNames(ctx context.Context) ([]string, error)
	GetAllTableRatings(ctx context.Context) ([]RatingsModel, error)
	InsertTableVotes(ctx context.Context, votesModel *VotesModel) error
	InsertTableRatings(ctx context.Context, ratings ...RatingsModel) error
}

type projectRepository struct {
	sqlDatabase *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{sqlDatabase: db}
}

func (r *projectRepository) GetAllUserNames(ctx context.Context) ([]string, error) {
	rows, err := r.sqlDatabase.QueryContext(ctx, getDistinctUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *projectRepository) GetAllTableRatings(ctx context.Context) ([]RatingsModel, error) {
	rows, err := r.sqlDatabase.QueryContext(
		ctx,
		getTableRatingsQuery,
	)
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

func (r *projectRepository) InsertTableVotes(ctx context.Context, votesModel *VotesModel) error {
	return r.sqlDatabase.QueryRowContext(
		ctx,
		insertTableVotesQuery,
		votesModel.UserName,
		votesModel.ImageWinner,
		votesModel.ImageLoser,
		votesModel.EloWinnerPrevious,
		votesModel.EloWinnerNew,
		votesModel.EloLoserPrevious,
		votesModel.EloLoserNew,
	).Scan(&votesModel.CreatedAt)
}

func (r *projectRepository) InsertTableRatings(ctx context.Context, ratings ...RatingsModel) error {
	tx, err := r.sqlDatabase.BeginTx(ctx, nil)
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
