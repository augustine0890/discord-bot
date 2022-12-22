FROM golang:1.18-alpine AS build

WORKDIR /build
COPY . .
# COPY prod.env .
RUN go mod download
RUN go build -o bot ./cmd/bot/main.go

#---------------------------------------

FROM alpine:3 AS final

WORKDIR /app
COPY --from=build /build/bot ./bot
ENTRYPOINT ["./bot"]