include ../tools/golangci-lint/rules.mk

develop:
	go run ./cmd/server/main.go

deploy:
	gcloud builds submit --tag gcr.io/senzr-313218/senzr
	gcloud run services update senzr-api --image=gcr.io/senzr-313218/senzr:latest --region=europe-west1

install:
	go mod tidy

test:
	go test ./...

lint: golangci-lint
