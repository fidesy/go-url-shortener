FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/go-url-shortener

FROM scratch

WORKDIR /app
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/main /usr/bin/

ENTRYPOINT ["main"]
CMD [ "main" ]