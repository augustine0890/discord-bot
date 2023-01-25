FROM golang:1.18-alpine AS build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o bot ./cmd/bot/main.go

#---------------------------------------

FROM alpine:3 AS final

WORKDIR /app
COPY --from=build /usr/local/go/ /usr/local/go/
COPY --from=build /build/bot ./bot
COPY *.env /app

ENV GOPATH="/usr/local/go/bin:${GOPATH}"
ENTRYPOINT ["./bot"]