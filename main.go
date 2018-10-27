package main

import (
	"github.com/fatih/color"
	"github.com/kataras/iris"
)

var (
	red       = color.New(color.FgRed).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	blue      = color.New(color.FgBlue).SprintFunc()
	magenta   = color.New(color.FgMagenta).SprintFunc()
	cyan      = color.New(color.FgCyan).SprintFunc()
	white     = color.New(color.FgWhite).SprintFunc()
	info      = color.New(color.FgWhite, color.BgGreen).SprintFunc()
	bgMagenta = color.New(color.FgWhite, color.BgMagenta).SprintFunc()
	bgYellow  = color.New(color.FgWhite, color.BgYellow).SprintFunc()
)

func main() {
	server := Server{}
	server.init()

	app := iris.Default()

	// set the view engine target to ./templates folder
	app.RegisterView(iris.HTML("./assets/templates", ".html").Reload(true))

	// Test serve html templates
	app.Get("/", server.home)

	// Serve the launch data
	app.Get("/json", server.homeJSON)

	app.Get("/update", server.update)

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
