package main

import "github.com/DinnieJ/njav-downloader/cmd"

var (
	Version  string
	AppName  string
	BuildSrc string
)

func main() {
	app := cmd.NewApplication(Version, AppName, BuildSrc)
	app.Start()
}
