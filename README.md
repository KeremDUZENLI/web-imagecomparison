# Image Comparison Website
https://keremduzenli.github.io/web-imagecomparison/


# BACKEND SUMMARY
```
    Database -> Run in Docker Container (PostgreSQL)
    Backend  -> Run Locally with Golang
```

## DOCKER (PostgreSQL)
```
    docker-compose build --no-cache
    docker-compose up
    docker ps
```

## BACKEND (Onion Architecture)
```
    /app/controller   -> Handles HTTP logic
    /app/middleware   -> HTTP Logs
    /app/model        -> Data structures
    /app/repository   -> DB operations
    /app/router       -> Sets routes
    /app/service      -> Business logic
    /database/connect -> DB Connection String
    /env/constants    -> Default Rating & KFactor
    /env/env.go       -> Load .env Variables
    /utils/shutdown   -> Shutdows Router Gracefully
    /main.go          -> App setup and router
```

# FRONTEND SUMMARY

## Generate images.json
```
npm install
npm run generate-images
```
