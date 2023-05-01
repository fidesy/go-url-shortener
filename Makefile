include .env
export 

test:
	go clean -testcache; go test -v ./...

rundb:
	docker run --name urlsdb -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=urls -p 5432:5432 -d postgres 

run-mongo:
	docker run --name urls-mongo -e MONGO_INITDB_ROOT_USERNAME=mongo -e MONGO_INITDB_ROOT_PASSWORD=mongo -dp 27017:27017 mongo

connect:
	docker exec -it urlsdb bash -c "psql -U postgres"

build:
	docker build --tag fidesy/go-url-shortener .

run:
	docker run --name go-url-shortener -dp 80:80 fidesy/go-url-shortener

migrate-up:
	migrate -source file:migrations -database 'postgres://postgres:postgres@localhost?sslmode=disable' -verbose up

migrate-down:
	migrate -source file:migrations -database 'postgres://postgres:postgres@localhost?sslmode=disable' -verbose down`

remove:
	docker rm -f urlsdb
	docker rm -f go-url-shortener
	docker rmi fidesy/go-url-shortener