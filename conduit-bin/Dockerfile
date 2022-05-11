FROM golang:1.18-alpine as build
EXPOSE 80

WORKDIR /app

COPY ./conduit-api ./conduit-api
COPY ./conduit-bin ./conduit-bin
COPY ./conduit-core ./conduit-core
COPY ./conduit-domain ./conduit-domain
COPY ./conduit-ent-gen ./conduit-ent-gen
COPY ./conduit-shared ./conduit-shared
COPY ./.env ./.env

COPY go.* ./
RUN go work sync

# Since we're using some libs that rely on gcc, add build-base before running building our go binary
RUN apk add --no-cache build-base && go build -a -o conduit ./conduit-bin

CMD ["./conduit", "-env", "docker", "-port", "8080"]