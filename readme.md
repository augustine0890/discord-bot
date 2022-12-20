# Discord Bot Application

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
- `docker build -t stats-bot .`
- Create and start containers
  - `docker compose up`
- Stop service: `docker compose stop`
- Stop and remove containers, networks: `docker-compose down --remove-orphans`
