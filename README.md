# DevZone API using Go

## Live Reloading using [Air](https://github.com/cosmtrek/air)

```shell
go install github.com/cosmtrek/air@latest

# Run app with live reload
$ docker-compose up -d devzone-db
$ air

# Run tests
$ go test -v ./...
```
