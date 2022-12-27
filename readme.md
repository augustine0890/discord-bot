# Discord Bot Application

## How to run the Bot
- `go run cmd/bot/main.go`
- Using the `-stage` flag
  - `go run cmd/bot/main.go -stage="dev"`
  - `go run cmd/bot/main.go -stage dev`
- Run with `go build`
  - `go build -o bot ./cmd/bot/main.go`
  - Build: `./bot -stage dev` --> development stage
## Install Docker and Go
- Update the installed packages
  - `sudo apt-get update`
- Remove Docker file
  - `sudo apt-get remove docker docker-engine docker.io`
- Install Docker
  - `sudo apt install docker.io`
- Install all the dependency packages
  - `sudo snap install docker`
  - Check Docker: `docker --version`
- Start the Docker service
  - `sudo service docker start`

## Build Docker image
- Build the image from Dockerfile
  - `docker build -t stats-bot:tags .`
- Run the new container from image
  - `docker run stats-bot:tags -stage dev`
- Remove Docker image:
  - `docker image rm -f image_id`
- Create and start containers:
  - `docker compose up --build`
  - `docker-compose --env-file prod.env up --build`
- Stop service: `docker compose stop`
- Stop and remove containers, networks: `docker-compose down --remove-orphans`
- Remove stopped containers, dangling images:
  - `docker system prune`