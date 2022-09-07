# typesense-docker
## Build Docker
```
docker-compose up --build
```

## Install package Golang
```
docker-compose exec golang go mod tidy
```

## Folder vendor
```
docker-compose exec golang go mod vendor
```

## Run Crawl
```
docker-compose exec golang go run main.go
```

## Website used typesense
https://techdaily.info/search?q=golang
