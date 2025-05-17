package app

import (
	"context"
	"database/sql"
)

const (
	CreateTableSurveysQuery = `
		CREATE TABLE IF NOT EXISTS surveys (
			username         TEXT 		  PRIMARY KEY,
			age              TEXT	 	  NOT NULL,
			gender           TEXT 		  NOT NULL,
			vr_experience    TEXT 		  NOT NULL,
			domain_expertise TEXT 		  NOT NULL,
			created_at    	 TIMESTAMPTZ  NOT NULL DEFAULT NOW()
		);`

	CreateTableVotesQuery = `
		CREATE TABLE IF NOT EXISTS votes (
			username 			 TEXT,
			image_winner 		 TEXT,
			image_loser 		 TEXT,
			elo_winner_previous  INTEGER,
			elo_winner_new 		 INTEGER,
			elo_loser_previous 	 INTEGER,
			elo_loser_new 		 INTEGER,
			created_at  		 TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);`

	CreateTableRatingsQuery = `
		CREATE TABLE IF NOT EXISTS ratings (
			image       TEXT 		 PRIMARY KEY,
			elo         INTEGER 	 NOT NULL,
			created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
			updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
		);`

	getDistinctUsersQuery = `
		SELECT DISTINCT username FROM votes;`

	getRatingsQuery = `
		SELECT image, elo, created_at, updated_at FROM ratings;`

	insertSurveyQuery = `
		INSERT INTO surveys (
			username, 
			age, 
			gender, 
			vr_experience, 
			domain_expertise
		)
		VALUES($1,$2,$3,$4,$5)
		ON CONFLICT (username) DO NOTHING;`

	insertVoteQuery = `
		INSERT INTO votes (
			username,
			image_winner,
			image_loser,
			elo_winner_previous,
			elo_winner_new,
			elo_loser_previous,
			elo_loser_new
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING created_at;`

	insertRatingQuery = `
		INSERT INTO ratings (image, elo)
			VALUES ($1, $2)
			ON CONFLICT (image) DO UPDATE SET 
				elo = EXCLUDED.elo,
				updated_at = NOW();`
)

type ProjectRepository interface {
	GetUsernames(ctx context.Context) ([]string, error)
	GetRatings(ctx context.Context) ([]RatingsModel, error)
	InsertSurvey(ctx context.Context, user SurveysModel) error
	InsertVote(ctx context.Context, votesModel *VotesModel) error
	InsertRating(ctx context.Context, ratings ...RatingsModel) error
}

type projectRepository struct {
	sqlDatabase *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{sqlDatabase: db}
}

func (r *projectRepository) GetUsernames(ctx context.Context) ([]string, error) {
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

func (r *projectRepository) GetRatings(ctx context.Context) ([]RatingsModel, error) {
	rows, err := r.sqlDatabase.QueryContext(ctx, getRatingsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []RatingsModel
	for rows.Next() {
		var m RatingsModel
		if err := rows.Scan(&m.Image, &m.Elo, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, nil
}

func (r *projectRepository) InsertSurvey(ctx context.Context, surveysModel SurveysModel) error {
	_, err := r.sqlDatabase.ExecContext(
		ctx,
		insertSurveyQuery,
		surveysModel.Username,
		surveysModel.Age,
		surveysModel.Gender,
		surveysModel.VRExperience,
		surveysModel.DomainExpertise,
	)
	return err
}

func (r *projectRepository) InsertVote(ctx context.Context, votesModel *VotesModel) error {
	return r.sqlDatabase.QueryRowContext(
		ctx,
		insertVoteQuery,
		votesModel.Username,
		votesModel.ImageWinner,
		votesModel.ImageLoser,
		votesModel.EloWinnerPrevious,
		votesModel.EloWinnerNew,
		votesModel.EloLoserPrevious,
		votesModel.EloLoserNew,
	).Scan(&votesModel.CreatedAt)
}

func (r *projectRepository) InsertRating(ctx context.Context, RatingsModels ...RatingsModel) error {
	tx, err := r.sqlDatabase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, ratingModel := range RatingsModels {
		if _, err := tx.ExecContext(
			ctx,
			insertRatingQuery,
			ratingModel.Image,
			ratingModel.Elo,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}
