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
- Add the ec2-user to the docker group so you can execute Docker #commands without using sudo.
  -  Create the docker group using groupadd command:
    - `sudo groupadd docker`
  - Add your user to this group with the usermod command:
    - `sudo usermod -aG docker $USER`
    - Verify that your user has been added to docker group: `groups`

## Build Docker image
- Build the image from Dockerfile
  - `docker build -t stats-bot:tags .`
- Run the new container from image (Run container in background)
  - `docker run -d --network="host" stats-bot:tags -stage dev`
  - `docker run -d --network="host" stats-bot:tags` (production stage)
- Remove Docker image:
  - `docker image rm -f image_id`
- Create and start containers:
  - `docker compose up --build`
  - `docker-compose --env-file prod.env up --build`
- Stop service: `docker compose stop`
- Stop and remove containers, networks: `docker-compose down --remove-orphans`
- Remove stopped containers, dangling images:
  - `docker system prune`
- List containers
  - `docker container ls -a` --> show all containers
## TODO
- Implement logging file
