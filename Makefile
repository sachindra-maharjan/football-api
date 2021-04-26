.PHONY: football-api
.PHONY: dbimport

football-api:
	@echo "Building the client binary"
	go build -o bin/football-api cmd/client/main.go

dbimport:
	@echo "Building the database import binary"
	go build -o bin/dbimport dbimport/cmd/main.go	
