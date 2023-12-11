# spotifyService

## Getting started
* start the server: `go run main.go` // Then follow the OAUTH2 process
* stop the server `ctrl+c`
* start the server again: `go run main.go`
* once there the docs of the API are here: `http://127.0.0.1:8000/docs/index.html`

## Testing

`go test -v ./...` or `make test`

# For Docs we use swagger
`go install github.com/swaggo/swag/cmd/swag@latest`

## update docs
`swag init --parseDependency`
