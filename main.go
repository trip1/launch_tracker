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

	// set the view engine target to ./templates folder
	app.RegisterView(iris.HTML("./assets/templates", ".html").Reload(true))

	// Test serve html templates
	app.Get("/", serveHTML)

	// Serve the launch data
	app.Get("/test", serveLaunches)

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}

func serveHTML(ctx iris.Context) {
	launches := fetchLaunches()
	ctx.ViewLayout("layout.html")
	fmt.Printf("%+v", launches.Data)
	ctx.ViewData("Launches", launches.Data)
	/*
		for k, v := range launches.Data {
			index := strconv.Itoa(k)
			key := "Launch" + index
			ctx.ViewData(key, v.Name)

			key = "Img" + index
			ctx.ViewData(key, v.Rocket.Image)

			key = "Date" + index
			ctx.ViewData(key, v.Start)
		}
	*/

	if err := ctx.View("index.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

// Serve launch data as json
func serveLaunches(ctx iris.Context) {
	ctx.JSON(fetchLaunches())
}

// Fetches next 10 launch
// from launchlibrary
func fetchLaunches() Launches {
	// Fetch data
	resp, err := http.Get("https://launchlibrary.net/1.4/launch/next/10")
	if err != nil {
		fmt.Println("Failed to get launch data")
	}

	// Close the body when finished
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	launches := Launches{}

	// Read resp into array
	err = json.Unmarshal(body, &launches)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v", launches)
	return launches
}
