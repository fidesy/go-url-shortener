# go-url-shortener
simple url shortener written in Go.

## Installation

1. Clone the repository.
```
git clone https://github.com/fidesy/go-url-shortener.git
cd go-url-shortener
```

2. Run app.
```
docker compose up -d
```

## Usage

Create short URL
```bash
curl -X POST -d '{"original_url": "https://vk.com"}' "http://localhost:8000/create"

>>> http://localhost:8000/ti2SMt
```

Get original URL and Redirect
```bash
curl http://localhost:8000/ti2SMt

>>> <a href="https://vk.com">Temporary Redirect</a>.
```
## Todo

1. User authorization (api keys)
2. Hash function for creating short url, now is using random sequence.
3. Add more tests