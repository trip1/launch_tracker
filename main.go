package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/", serveLaunches)
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}

func serveLaunches(ctx iris.Context) {
	ctx.Write([]byte("Hello"))
	fetchLaunches()
}

func fetchLaunches() {
	resp, err := http.Get("https://launchlibrary.net/1.4/launch/next/10")
	if err != nil {
		fmt.Println("Failed to get launch data")
	}

	decoder := json.NewDecoder(resp.Body)
	launches := [10]Launch{}

	err = decoder.Decode(&launches)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(launches)
}
