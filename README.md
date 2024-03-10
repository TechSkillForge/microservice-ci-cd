# Golang

This is a simple API in Go using Gin-Gonic with an in-memory data store and routes to handle Create and Read operations of Products, to demonstrate the use of CI/CD with GitHub Actions.

## Pre-requisites

- Docker

## How to Run

1. Build the Docker image

```bash
docker build -t go-products-msvc:local .
```

2. Run the Docker container

```bash
docker run -d -p 8080:8080 go-products-msvc:local
```

## How to Test

You can test the application routes using `curl`:

1. Create a Product

```bash
$ curl -X POST http://localhost:8080/products -d '{"name": "item1"}'

{
    "data": {
        "name": "item1"
    },
    "success": true
}
```

2. Get all Products

```sh
$ curl http://localhost:8080/products

{
    "data": [
        {
            "name": "item1"
        }
    ],
    "success": true
}
```
