FROM golang:1.23.4-alpine3.20

RUN apk update
RUN apk add --no-cache curl git
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

COPY ./ /goherent

WORKDIR /goherent
