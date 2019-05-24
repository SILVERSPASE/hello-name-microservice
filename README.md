# Hello name microservice
Here we have two microservices:
- greeter service which has business logic - gets name and returns greeting phrase with it.
- client service which listens port :8080 for requests and calls greeter service via rcp and pass name to it.

All code here is oversimplified to make everything clear

## How to spin up

Go to the root of the repo

Run in the first tab

```go run ./greeter-service/cmd/main.go```

Open the second tab and run next command

```go run ./client/cmd/main.go```

And open the third tab and send curl request to localhost:

```curl localhost:8080```

The answer should be `Hello, Wanderer`, because there is no name passed with request, client chose default one.

If you pass name - you receive phrase with that name

```curl localhost:8080 -d '{"name":"Alice"}'```

The answer should be `Hello, Alice`

## What is under the hood

With first command we started greeter-service with out business logic - it has rpc function which do simple work and do it well - gets name and returns the phrase.

The second command spined up client service which listens port 8080 for requests and calls greeter-service rpc function

By sending curl request we activate client chain of parsing body - for name param and conducting rpc call to greeter-service.

## TODO

- Dockerize it!

- Authorisation




