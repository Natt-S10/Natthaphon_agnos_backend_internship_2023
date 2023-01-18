# Natthaphon Agnos Aackend Internship 2023
This Repository is build for Agnos backend internship 2023 application. 

## Deployment Instruction

### 1. Setup PostgresSQL on Docker
> running `docker-compose.yml` in the repository root . 
>```shell
>$ docker-compose up
>```
PostgresSQL server is hosted on `localhost:5438` with username: `dev` and password `12345678` 

### 2. Setup and run the API server
> download dependencies: 
> ```shell
> $ go mod tidy
> ```
> run server
> ```shell
> $ go run ./main.go
> ```
The server is serving on `localhost:8080`

## API Routes
| use case | route | method | payload |
|----------|-------|--------|---------|
| Ping     | `/ping`| GET   |  none   |
| ==Request for strong password steps== | `/api/strong_password_steps` | POST| ```{"init_password": "aA1"}```

## Testing
Running all unit test using the command below
```shell
$ go test ./...
``` 
The test includes 
> - passwd_correction_test
> - password_util_test
