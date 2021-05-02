
image = senzr_db
db_name = senzr
db_password = gbgftw

all:

develop: run-database
	go run ./cmd/server/main.go

run-database:
	docker stop ${image}
	docker rm ${imag}
	docker run --name ${image} -e POSTGRES_PASSWORD=${db_password} -e POSTGRES_DB ${db_name} -d postgres
