# Go-health-api

A simple API service written in Golang.
Im building this project as proof of work in hopes of landing a DevOps / Cloud role. 

## How to run locally

You can either use `go run main.go` on the root directory or run the executable file product of running `go build`

## Endpoints

- **"/":** returns a json body with a simple message on success.

```bash
$ curl -X GET http://127.0.0.1:8080/
{"message": "Health API is up and running"}
```
- **"/health":** returns a 200 OK http status code and a simple json on success. Means the api is up and running.

```bash
$ curl -X GET http://127.0.0.1:8080/health
{"status": "ok"}
```

If you use another method like POST you will get a **405 Method Not Allowed**.

```bash
$ curl -X POST http://127.0.0.1:8080/health
Method Not Allowed
```
