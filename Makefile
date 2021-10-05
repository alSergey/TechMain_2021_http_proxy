PROJECT_DIR := ${CURDIR}

MAIN_BINARY=main

build:
	go build -o ${MAIN_BINARY} cmd/main.go

run-go:
	go run cmd/main.go

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo