all:

develop: run-database
	go run ./cmd/server/main.go

run-database:
	docker compose up -d

deploy:
	gcloud builds submit --tag gcr.io/senzr-313218/senzr
	gcloud run services update senzr-api --image=gcr.io/senzr-313218/senzr:latest --region=europe-west1

test:
	go test ./...