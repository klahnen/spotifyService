# spotifyService

## Getting the token
go run main.go // Then follow the OAUTH2 process

## Testing
export BEARER_TOKEN=<access token>
go test -v ./...

# swagger
go install github.com/swaggo/swag/cmd/swag@latest
