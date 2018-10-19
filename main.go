package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	ctx.JSON(fetchLaunches())
}

func fetchLaunches() Launches {
	resp, err := http.Get("https://launchlibrary.net/1.4/launch/next/10")
	if err != nil {
		fmt.Println("Failed to get launch data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	launches := Launches{}

	err = json.Unmarshal(body, &launches)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", launches)
	return launches
}
