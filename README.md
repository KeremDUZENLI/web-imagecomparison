# VR IMAGE COMPARISON SURVEY

**Static Website:** [https://keremduzenli.github.io/web-imagecomparison](https://keremduzenli.github.io/web-imagecomparison/)

**Live Website:** [https://tinyurl.com/vr-compare](https://web-imagecomparison-api.onrender.com/)


## OVERVIEW

A VR scene realism comparison tool that collects user preferences via an ELO-based ranking system.

| Feature                 | Description                                               |
|-------------------------|-----------------------------------------------------------|
| **User Survey**         | Collect demographic info (age, gender, VR experience, domain expertise) |
| **Pairwise Comparison** | Display random non-repeating image pairs                 |
| **ELO Ranking**         | Automatically adjust image scores                        |
| **Leaderboard**         | View top-ranked images after voting                       |
| **Session Persistence** | Handles page reloads and unfinished votes with sessionStorage |


## GETTING STARTED

### Prerequisites

* [Go](https://golang.org/doc/install) (v1.18+)
* [Docker](https://docs.docker.com/compose/install/)


### Configuration

Copy `.env.example` to `.env` in `backend/` and update:

```
DB_HOST     =localhost
DB_USER     =test
DB_PASSWORD =test
DB_NAME     =postgres_test
DB_PORT     =5430
DB_SSLMODE  =disable
SERVER_PORT =5501
```


### Installation

1. **Clone the repo**
```bash
git clone https://github.com/KeremDUZENLI/web-imagecomparison.git
cd web-imagecomparison/backend
```

2. **Start the database**
```bash
docker-compose build --no-cache
docker-compose up
docker ps
```

3. **Run the backend**
```bash
go run main.go
```

4. **Open the frontend**
```bash
http://localhost:5501
```


### Updating Images List

If you add or remove images in `images/`, update `_images.json`:

```bash
cd  scripts/
npm install
npm run createImageJSON
```


## PROJECT STRUCTURE

```
web-imagecomparison/
├── backend/                        # Go API and database setup

│   ├── app/                        # Core application layers
│   │   ├── controller.go           # HTTP handlers for surveys, votes, users, ratings
│   │   ├── middleware.go           # Request logging and static cache control
│   │   ├── model.go                # Data models (SurveysModel, VotesModel, RatingsModel)
│   │   ├── repository.go           # Database queries and upserts
│   │   ├── router.go               # Route registration and HTTP server mux
│   │   └── service.go              # Business logic (ELO calculation, workflows)
│   │
│   ├── database/                   # DB connection
│   │   └── connect.go              # Postgres connection helper (sslmode, ping)
│   │
│   ├── env/                        # Environment config
│   │   ├── constants.go            # DEFAULT_RATING, K_FACTOR
│   │   └── env.go                  # Load and validate .env variables
│   │
│   ├── utils/                      # Auxiliary tools
│   │   └── shutdown.go             # Graceful shutdown on SIGINT/SIGTERM
│   │
│   ├── .env                        # Local environment variables (DB, SSL, SERVER_PORT)
│   ├── .env.example                # Template for env settings
│   ├── docker-compose.yml          # Local Postgres container definition
│   ├── go.mod                      # Go module definition
│   ├── go.sum                      # Module checksums
│   └── main.go                     # Entry point: load config, init DB, start server
│
├── css/                            # Global stylesheet (styles.css)
│   └── styles.css
│
├── images/                         # VR scene image assets
│   └── ... (e.g. 1.jpg, 2.jpg, etc.)
│
├── js/                             # Frontend logic and UI modules
│   ├── core/                       # Session logic
│   │   └── matchSession.js         # Random non‑repeating pair generation
│   │
│   ├── env/                        # Frontend constants
│   │   └── constants.js            # MIN_VOTES, TOPN
│   │
│   ├── infrastructure/             # API fetch wrappers
│   │   ├── getRatings.js
│   │   ├── getUsernames.js
│   │   ├── postSurvey.js
│   │   └── postVote.js
│   │
│   ├── ui/                         # DOM update helpers
│   │   ├── loadImages.js           # Load images/_images.json
│   │   ├── setText.js              # Set textContent by ID
│   │   ├── showLeaderboard.js      # Render top‑N leaderboard
│   │   └── showPair.js             # Display A/B image pair and progress
│   │
│   ├── utils/                      # Utility helpers
│   │   └── waitForEnterKey.js      # Promise for Enter key event
│   │
│   ├── compare.js                  # Orchestrates compare page logic
│   ├── finish.js                   # Handles finish page and cached leaderboard
│   └── index.js                    # Handles survey intro form submission
│
├── scripts/                        # (Optional) build/generation scripts
│   └── generate-images.js          # Script to create images/_images.json
│
├── compare.html                    # Voting page template
├── finish.html                     # Thank‑you & leaderboard page template
├── index.html                      # Survey intro page template
├── LICENSE                         # BSD 3‑Clause License
└── README.md                       # This documentation
```


## BACKEND (Onion Architecture)

GO:
```
app/controller   --> HTTP handlers
app/middleware   --> Logging & cache control
app/model        --> Data structures
app/repository   --> Database queries
app/service      --> Business logic (ELO computation)
app/router       --> Route registration
database/connect --> Postgres connection
env              --> Env var loading & defaults
utils/shutdown   --> Graceful server shutdown
main.go          --> Application entry point
```


## FRONTEND (Responsive)

CSS:
```
css/style          --> Responsive, accessible design
```

Vanilla JS:
```
/core/matchSession --> Class
/env/constants     --> Constant values
/infrastructure/   --> API calls
/ui/               --> UI rendering
/utils/            --> Extra tools
/                  --> Session logics
```


## LICENCE

This project is released under the [BSD 3-Clause License](LICENSE).


## DISCLAIMER

This repository is intended **only for educational and research purposes**.


## SUPPORT MY PROJECTS

If you find this resource valuable and would like to help support my education and doctoral research, please consider treating me to a cup of coffee (or tea) via Revolut.

<div align="center">
  <a href="https://revolut.me/krmdznl" target="_blank">
    <img src="https://img.shields.io/badge/Support%20My%20Projects-Donate%20via%20Revolut-orange?style=for-the-badge" alt="Support my projects via Revolut" />
  </a>
</div> <br>
