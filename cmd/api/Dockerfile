FROM golang:1.18-alpine as build
EXPOSE 80

WORKDIR /app

COPY go.* ./
RUN go mod download
RUN go mod verify

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./ent ./ent
COPY ./.env ./.env

RUN go build -o conduit ./cmd/api

ENTRYPOINT ["/api"]
CMD ["-env", "docker", "-port", "8080"]