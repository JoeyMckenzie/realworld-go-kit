FROM golang:1.18-alpine as build
EXPOSE 80

WORKDIR /app

COPY ./src/conduit-api ./conduit-api
COPY ./src/conduit-bin ./conduit-bin
COPY ./src/conduit-core ./conduit-core
COPY ./src/conduit-domain ./conduit-domain
COPY ./src/conduit-ent-gen ./conduit-ent-gen
COPY ./src/conduit-shared ./conduit-shared
COPY ./.env ./.env

COPY go.docker.work ./go.work
COPY go.work.sum ./
RUN go work sync

# Since we're using some libs that rely on gcc, add build-base before running building our go binary
RUN apk add --no-cache build-base && go build -a -o conduit ./conduit-bin

CMD ["./conduit", "-env", "docker", "-port", "8080", "-seed", "true"]
