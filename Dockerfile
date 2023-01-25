FROM golang:1.18-alpine AS build

WORKDIR /build
COPY . .
# install cron
RUN apt-get update && apt-get install cron -y -qq
RUN go mod download
RUN go build -o bot ./cmd/bot/main.go

#---------------------------------------

FROM alpine:3 AS final

WORKDIR /app
COPY --from=build /build/bot ./bot
COPY *.env /app
ENTRYPOINT ["cron", "-f", "./bot"]