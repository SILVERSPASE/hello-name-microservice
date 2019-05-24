# Hello %name% microservice
Here we have two microservices:
- greeter service which has business logic - gets name and returns greeting phrase with it.
- client service which listens port :8080 for requests and calls greeter service via rcp and pass name to it.

## How to spin up
go to the root of the repo
run in the first tab
```go run ./greeter-service/cmd/main.go```
open the second tab and run next command
```go run ./client/cmd/main.go```
and open the third tab and send curl request to localhost:
```curl localhost:8080```
The answer should be `Hello, Wandered`


