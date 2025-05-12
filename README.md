# web-imagecomparison
Image Comparison Algorithm
https://keremduzenli.github.io/web-imagecomparison/

## docker
docker-compose build --no-cache
docker-compose up
docker ps

## npm
cd .\backend\
npm install
npm start


# Setup Summary
✅ Backend    : Run locally with node index.js
✅ PostgreSQL : Run in Docker container
❌ Container  : No Docker for backend

## Clean Archicture (Backend)
/common     -> Commons and Tools
/controller -> Handles HTTP logic
/database   -> PostgreSQL
/model      -> Data structures
/repository -> DB operations
/service    -> Business logic
/main.go    -> App setup and router
