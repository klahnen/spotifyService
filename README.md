# spotifyService

## Getting the token
* start the server: `go run main.go` // Then follow the OAUTH2 process
* stop the server `ctrl+c`
* start the server again: `go run main.go`

## Testing

go test -v ./...

# swagger
go install github.com/swaggo/swag/cmd/swag@latest

## update docs
swag init --parseDependency
