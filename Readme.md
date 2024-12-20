# Go Assessment Project

This project was created as part of a job application assessment.
## Description
This project consists of three separate services.

- Service1 and Service2 each expose a GetData method accessible via gRPC. When a request is made, a number between 1 and 3 should be provided to receive a corresponding simple string response. If the number is outside this range, a "not found" response will be returned.
- Service3 provides a GetResult method. This method first calls GetData from Service1 while simultaneously inserting random data into a SQLite database. If both operations succeed without errors, it proceeds to call GetData from Service2 and prints the result.

On the database side, SQLite is used for its simplicity, and GORM is utilized for enhanced readability and convenience in database operations.
## Dependencies

```bash
go get -u gorm.io/gorm
```
```bash
go get github.com/mattn/go-sqlite3
```
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Executing program

In Service1 and Service2, the main function located in the server directory must be executed to expose the GetData methods on ports 8080 and 8081, respectively.

In Service3, the main function located in the cmd directory should then be executed to call these methods and perform the query operation.