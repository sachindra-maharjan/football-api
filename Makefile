.PHONY: football-api

football-api:
	@echo "Building the client binary"
	go build -o bin/football-api cmd/client/main.go