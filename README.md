# Golang Challenge 5 Final Project

[API Specs](https://gospecs.monstercode.net)

Seeting env
```bash 
# copy file .env.example
cp .env.example .env
```

Running Project
```bash
# running all service using docker-compose
docker-compose up -d --force-recreate
# running final project
docker-compose exec exec go go run .
# project running in port 8000
http://localhost:8000/api/v1
```

Create Dokumentation Swagger
```bash 
# swagger init
swag init --parseDependency
# swagger documentation running in port 8080
http://localhost:8080/
```

Check clean code using goreportcard
```bash
goreportcard-cli -v
```

Running test:
```bash
docker-compose up test
# or
go test -v -coverprofile cover.txt ./...
```

Get Coverage:
```bash
go tool cover -html cover.txt
# or
go tool cover -func cover.txt
```