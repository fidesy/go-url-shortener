# go-url-shortener
simple url shortener written in Go.

## Installation

1. Clone the repository.
```
git clone https://github.com/fidesy/go-url-shortener.git
cd go-url-shortener
```

2. Rename .env.example to .env
```
cp .env.example .env
``` 

3. Select preferable database in [./configs/config.yaml](./configs/config.yaml#3) 
field *database* options: postgres, mongo

4. Run app.
```
docker compose up -d
```

## Usage

Sign up
```bash
curl -X POST -d '{"name": "John", "username": "john", "password": "john"}' http://localhost:8000/auth/sign-up

# {"id": 1}
```

Sign in and get authorization token
```bash
curl -X POST -d '{"username": "john", "password": "john"}' http://localhost:8000/auth/sign-in

# {"token": "YOUR_TOKEN"}
```
Create short URL
```bash
curl -X POST -H "Authorization: Bearer <YOUR_TOKEN>" -d '{"original_url": "https://vk.com"}' "http://localhost:8000/create"

# {"short_url": "http://localhost:8000/ti2SMt"}
```

Get original URL and Redirect
```bash
curl http://localhost:8000/ti2SMt

# <a href="https://vk.com">Temporary Redirect</a>.
```