package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Vote struct {
	ID                int       `json:"id,omitempty"`
	UserName          string    `json:"userName"`
	ImageA            string    `json:"imageA"`
	ImageB            string    `json:"imageB"`
	ImageWinner       string    `json:"imageWinner"`
	ImageLoser        string    `json:"imageLoser"`
	EloWinnerPrevious int       `json:"eloWinnerPrevious"`
	EloWinnerNew      int       `json:"eloWinnerNew"`
	EloLoserPrevious  int       `json:"eloLoserPrevious"`
	EloLoserNew       int       `json:"eloLoserNew"`
	CreatedAt         time.Time `json:"createdAt,omitempty"`
}

func connectDB() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "postgres")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to DB: %w", err)
	}
	log.Println("✅ Connected to PostgreSQL")
	return db, nil
}

func initTable(db *sql.DB) error {
	query := `
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
        );
    `
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	log.Println("✅ Table 'votes' is ready")
	return nil
}

func makeVotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📩 Received /api/votes POST request")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			log.Printf("Incoming %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
			return
		}

		var v Vote
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Decoded Vote %#v", v)
		// Insert vote
		query := `
			INSERT INTO votes (
				user_name, image_a, image_b, image_winner, image_loser,
				elo_winner_previous, elo_winner_new, elo_loser_previous, elo_loser_new
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			RETURNING id, created_at`
		var id int
		var createdAt time.Time
		err := db.QueryRow(
			query,
			v.UserName, v.ImageA, v.ImageB, v.ImageWinner, v.ImageLoser,
			v.EloWinnerPrevious, v.EloWinnerNew, v.EloLoserPrevious, v.EloLoserNew,
		).Scan(&id, &createdAt)
		if err != nil {
			// Print the full SQL error so we know what’s wrong
			log.Printf("Insert error: %T %v", err, err)
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		v.ID = id
		v.CreatedAt = createdAt
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(v)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	_ = godotenv.Load()

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initTable(db); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir("../docs")))
	http.HandleFunc("/api/votes", makeVotesHandler(db))

	port := getEnv("PORT", "8080")
	log.Printf("✅ Server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
