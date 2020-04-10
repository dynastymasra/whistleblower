# Whistleblower

[![Go](https://img.shields.io/badge/go-1.14.0-00E5E6.svg)](https://golang.org/)
[![Docker](https://img.shields.io/badge/docker-19.03-2885E4.svg)](https://www.docker.com/)
[![Postgres](https://img.shields.io/badge/postgres-12.2-27527D.svg)](https://www.postgresql.org/)

WhistleNews is a news portal for whistleblowers. Registered authors would be able to post articles without the
need to reveal their identity publically.

## Architecture

This project try to implement [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) 
and the Apps architecture can use [Kong](https://konghq.com/kong/) for API Gateway and can use [Kong Authentication](https://docs.konghq.com/hub/)
for authentication mechanism. Right now this service only use server key to access the endpoints.

Event Driven architecture use message broker (RabbitMQ, Kafka, NATS, etc).  

## Libraries

Use [Go Module](https://blog.golang.org/using-go-modules) for install all dependencies required this application.

## How To Run and Deploy

Before run this service. Make sure all requirements dependencies has been installed likes **Golang, Docker, and Postgres**

### Local

Use command go ```go run main.go``` in root folder for run this application.

- ```go run main.go migrate:run``` Command to run database migration
- ```go run main.go migrate:rollback``` Command to rollback database migration
- ```go run main.go migrate:create``` Command to create database migration file

### Docker

**whistleblower** uses docker multi stages build, minimal docker version is **17.05**. If docker already installed use command.

This command will build the images.
```bash
docker build -f Dockerfile -t $(IMAGE):$(VERSION) .
```

Run this command for database migration
```bash
docker-compose run --rm $(NAME) migrate:run
```

To run service use this command
```bash
docker run --name whistleblower -d -e ADDRESS=:8080 -e <environment> $(IMAGE):$(VERSION)
```

### Shell Script

Make sure *docker-compose* already installed. to run this script.

Use this command for the first time. this script will run docker compose and run the database migration.
```bash
./start.sh run
```

Use this command to run application without database migration. Need to run ```./start.sh run``` before use this command
```bash
./start.sh up
```

## Test

For run unit test, from root project you can go to folder or package and execute command, this project have two kinds of test, 
Integration Test and Unit Test. 
```bash
go test -v -cover -coverprofile=coverage.out -covermode=set
go tool cover -html=coverage.out
```
`go tool` will generate GUI for test coverage. Available package or folder can be tested.

+ **Integration Test**
    - `/infrastructure/web/handler`
    - `/article`
    - `/viewer`
+ **Unit Test**
    - `/article/handler/http`
    - `/viewer/handler/http`
    
## API Documentation

This service documentation uses [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/0b751a192febb1c370d8).


## Environment Variables

+ `SERVER_PORT` - Address application is used default is `:8080`
+ `LOG_FORMAT` - Format specific for log, default value is `text`
  - `text` - Log format will become standard text output, this used for development
  - `json` - Log format will become *JSON* format, usually used for production
+ `LOG_LEVEL` - Log level default is `debug`
+ `POSTGRES_ADDRESS` - Database hostname
+ `POSTGRES_DATABASE` - Database name
+ `POSTGRES_USERNAME` - Database username
+ `POSTGRES_PASSWORD` - Database Password
+ `POSTGRES_LOG_ENABLED` - Database log enabled, value `true` or `false`
+ `POSTGRES_MAX_OPEN_CONN` - Database max open connection
+ `POSTGRES_MAX_IDLE_CONN` - Database max idle connection