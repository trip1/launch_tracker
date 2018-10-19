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

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
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
