all:

develop: run-database
	go run ./cmd/server/main.go

run-database:
	docker stop senzr_db
	docker rm senzr_db
	docker run --name senzr_db -e POSTGRES_PASSWORD=gbgftw -d postgres
	