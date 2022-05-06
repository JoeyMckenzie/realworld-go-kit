FROM golang:1.18-alpine as build
EXPOSE 80

WORKDIR /app

COPY go.* ./
RUN go mod download
RUN go mod verify

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./ent ./ent
COPY ./.env ./.env

# Since we're using some libs that rely on gcc, add build-base before running building our go binary
RUN apk add --no-cache build-base && go build -a -o conduit ./cmd/api

CMD ["./conduit", "-env", "docker", "-port", "8080"]