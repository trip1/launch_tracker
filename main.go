package main

import (
	"github.com/kataras/iris"
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
