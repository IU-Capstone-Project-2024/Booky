# Include variables from the .env file
include .env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/booky
run/booky:
	go run ./cmd/booky -fileStorage=S3

.PHONY: generate/booky
generate/booky:
	@echo 'Generating proto files...'
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	api/booky/booky.proto

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies, format, vet and test all code
.PHONY: audit
audit:
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/booky application
.PHONY: build/booky
build/booky:
	@echo 'Building cmd/booky...'
	go build -o=./bin/booky ./cmd/booky -fileStorage=S3
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/booky ./cmd/booky -fileStorage=S3