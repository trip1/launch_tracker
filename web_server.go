package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

// Server handles HTTP traffic
// and state management of data
type Server struct {
	Launches        Launches      `json:"launches"`
	UpdatedAt       time.Time     `json:"updated_at"`
	TimeSinceUpdate time.Duration `json:"time_since_update"`
}

func (s *Server) init() {
	s.fetchLaunches()
	fmt.Println("Init complete")
}

func (s *Server) fetchLaunches() {
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

	s.Launches = launches
}

func (s *Server) home(ctx iris.Context) {
	ctx.ViewLayout("layout.html")
	ctx.ViewData("Launches", s.Launches.Data)

	if err := ctx.View("index.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}

func (s *Server) homeJSON(ctx iris.Context) {
	ctx.JSON(s.Launches)
}
