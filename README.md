# Web Image Comparison

A beginner-friendly web app that lets users compare two images and ranks them using an Elo algorithm.

**Live Demo:** [https://keremduzenli.github.io/web-imagecomparison/](https://keremduzenli.github.io/web-imagecomparison/)

---

## Prerequisites

* Go (1.18+)
* Node.js & npm (v16+)
* Docker & Docker Compose (for PostgreSQL)

---

## Quick Start

### 1. Clone the repo

```bash
git clone https://github.com/keremduzenli/web-imagecomparison.git
cd web-imagecomparison
```

### 2. Configure

Copy `.env.example` to `.env` and enter your database credentials and `SERVER_PORT`.

### 3. Start PostgreSQL (Docker)

```bash
docker-compose up -d database
```

### 4. Run Backend

```bash
cd backend
go run main.go
```

### 5. Serve Frontend

```bash
cd docs
npx http-server .
# then open http://localhost:8080
```

---

## Project Structure

```
/web-imagecomparison
├── backend    # Go server code
├── docs       # Static frontend (HTML, CSS, JS)
├── docker-compose.yml
└── .env*      # config file
```

---

## API: Record a Vote

**POST** `/api/votes`

Request JSON:

```json
{
  "userName": "Alice",
  "imageA": "a.jpg",
  "imageB": "b.jpg",
  "imageWinner": "a.jpg",
  "imageLoser": "b.jpg",
  "eloWinnerPrevious": 1500,
  "eloWinnerNew": 1516,
  "eloLoserPrevious": 1500,
  "eloLoserNew": 1484
}
```

Response: saved vote with `id` and `createdAt`.

---

## License

MIT

\*Files ending in `*` should be created from their `.example` counterparts.


✅ STRUCTURE OVERVIEW (Current Mapping)
Layer	              Responsibility	                       Current File(s)
Presentation	      HTTP handlers (controllers)	           controller.go, router.go
Application	        Business logic, use cases	             service.go
Domain	            Models/entities	                       model.go
Infrastructure	    Database access, external systems	     repository.go, connect.go
Startup	            App init, routing, environment	       main.go, app.go, env.go
