.DEFAULT_GOAL := runWithLocalCmd

.PHONY: prerequisites fmt vet build test runWithLocalCmd setupGitHooks

fmt: prerequisites
	go fmt ./...

vet: prerequisites fmt
	go vet ./...

build: prerequisites vet
	go build

test: prerequisites
	go test ./...

runWithLocalCmd: prerequisites
	go run main.go ./...

prerequisites:
	cp .githooks/pre-push .git/hooks/pre-push
