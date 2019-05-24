# Hello %name% microservice
Here we have two microservices:
- greeter service which has business logic - gets name and returns greeting phrase with it.
- client service which listens port :8080 for requests and calls greeter service via rcp and pass name to it.

## How to spin up

Go to the root of the repo

Run in the first tab

```go run ./greeter-service/cmd/main.go```

Open the second tab and run next command

```go run ./client/cmd/main.go```

And open the third tab and send curl request to localhost:

```curl localhost:8080```

The answer should be `Hello, Wanderer`

If you pass name - you receive phrase with that name

```curl localhost:8080 -d '{"name":"Alice"}'```

The answer should be `Hello, Alice`


