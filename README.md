# go-url-shortener
simple url shortener written in Go.

## Installation

1. Install postgres docker image and run it
```bash
docker pull postgres

docker run --name urlsdb \
   -e POSTGRES_USER=xsecretuser \
   -e POSTGRES_PASSWORD=kj1890Opokb19lf \
   -e POSTGRES_DB=urls \
   -dp 5432:5432 postgres 
```
2. Install url-shortener app and run it
```bash
docker pull fidesy/go-url-shortener

docker run --name go-url-shortener \
    -dp 80:80 fidesy/go-url-shortener
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