FROM golang:1.23.4-alpine3.20

RUN apk update
RUN apk add --no-cache curl git openssh nodejs npm
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
RUN npm install -g @anthropic-ai/claude-code

COPY ./ /goherent

WORKDIR /goherent
