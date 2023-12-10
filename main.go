package main

import (
	"github.com/klahnen/spotifyService/app"
)

// @title          spotifyService API
// @version         1.0
// @description     This is a sample server

// @contact.name   Fernando
// @contact.email  klahnen@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8000
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app := app.App{}
	app.Initialize()
	app.Run()
}
