FROM golang:1-bullseye AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

# COPY . .
COPY ./cmd ./internal ./
RUN go build -v -o /usr/local/bin/app ./main.go

# Stage two - we'll utilize a second container to run our built binary from our first container - slim containers!
FROM debian:bullseye-slim as deploy

# Let's install all the necessary runtime tools on the container
RUN set -eux; \
    export DEBIAN_FRONTEND=noninteractive; \
    apt update; \
    apt install -y --no-install-recommends bind9-dnsutils iputils-ping iproute2 curl ca-certificates htop; \
    apt clean autoclean; \
    apt autoremove -y; \
    rm -rf /var/lib/{apt,dpkg,cache,log}/;

# Let's work from a self contained directory for all of our deployment needs
WORKDIR /deploy

# We need the artifact from the build container, so let's grab it
COPY --from=build /usr/local/bin/app ./

# Let's expose port 80 as we'll need fly's internal port mapping also assumes 80
EXPOSE 80
EXPOSE 443

# Finally, boot up the API
CMD ["./app"]