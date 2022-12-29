include .env
export 

test:
	go test -v ./...

rundb:
	docker run --name urlsdb -e POSTGRES_USER=xsecretuser -e POSTGRES_PASSWORD=kj1890Opokb19lf -e POSTGRES_DB=urls -p 5432:5432 -d postgres 

connect:
	docker exec -it urlsdb bash -c "psql -U xsecretuser -d urls"

build:
	docker build --tag fidesy/go-url-shortener .

run:
	docker run --name go-url-shortener -dp 80:80 fidesy/go-url-shortener

remove:
	docker rm -f urlsdb
	docker rm -f go-url-shortener
	docker rmi fidesy/go-url-shortener