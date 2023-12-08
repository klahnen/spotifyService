package main

import "github.com/klahnen/spotifyService/app"

func main() {
	app := app.App{}
	app.Initialize()
	app.Run()
}
