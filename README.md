# go-url-shortener
simple url shortener written in Go.

## Installation

1. Clone the repository.
```
git clone https://github.com/fidesy/go-url-shortener.git
cd go-url-shortener
```
2. Create .env file with the following variables. The HOST variable is only needed to create a response with a short URL.
```
# Example
HOST=http://localhost
PORT=80

# Postgres credentials to create the database
POSTGRES_USER=xsecretuser
POSTGRES_PASSWORD=kj1890Opokb19lf
POSTGRES_DB=urls

# We use urlsdb host - PostgreSQL docker container name 
DB_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@urlsdb/${POSTGRES_DB}?sslmode=disable
DB_NAME=postgresql
```
3. Run app.
```
docker compose up -d
```
## Usage

Create short URL
```bash
curl -X POST "http://localhost/create?url=https://vk.com"

>>> http://localhost:80/ti2SMt
```

Get original URL and Redirect
```bash
curl http://localhost/ti2SMt

>>> <a href="https://vk.com">Permanent Redirect</a>.
```
## Todo

1. User authorization (api keys)
2. Hash function for creating short url, now is using random sequence.
3. Add more tests