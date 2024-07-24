# stock-analyzer

## Execute on local machine

> go build cmd/api/main.go

> ./main

## Execute single container

> docker build -f Dockerfile -t server-go .

> docker run -p 127.0.0.1:3000:3000 server-go

## With docker-compose

> docker-compose build
> docker-compose up
