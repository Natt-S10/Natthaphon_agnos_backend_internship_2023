# Natthaphon Agnos Aackend Internship 2023
This Repository is build for Agnos backend internship 2023 application. 

## Deployment Instruction
At the repo's root directory, run the following command to starts both server and SQL server
```shell
$ docker-compose up -d
```
- PostgresSQL server is hosted on `localhost:5438` with username: `dev` and password `12345678` 
- The server is serving on `localhost:8080`

---
## API Routes
| use case | route | method | payload |
|----------|-------|--------|---------|
| Ping     | `/ping`| GET   |  none   |
| __Request for strong password steps__ | `/api/strong_password_steps` | POST| ```{"init_password": "aA1"}```
---
## Testing
Running all unit test using the command below
```shell
$ go test ./...
``` 
The test includes 
> - passwd_correction_test
> - password_util_test
