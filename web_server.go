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
	Launches   Launches    `json:"launches"`
	Ticker     time.Ticker `json:"ticker"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Update     chan bool   `json:"update"`
	StopUpdate chan bool   `json:"stop"`
}

func (s *Server) init() {
	s.fetchLaunches()

	s.Update = make(chan bool)
	s.StopUpdate = make(chan bool)

	fmt.Println("Init complete")
	go s.interval()
}

func (s *Server) fetchLaunches() {
	resp, err := http.Get("https://launchlibrary.net/1.4/launch/next/10&mode=verbose")
	if err != nil {
		fmt.Println("Failed to get launch data")
	}

	// Close the body when finished
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// Read resp into array
	err = json.Unmarshal(body, &s.Launches)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) interval() {
	ticker := time.NewTicker(6 * time.Hour)

	for {
		select {
		case <-s.Update:
			fmt.Println("Update")
			s.fetchLaunches()
			s.UpdatedAt = time.Now()
		case <-ticker.C:
			s.fetchLaunches()
			s.UpdatedAt = time.Now()
		}
	}
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

func (s *Server) update(ctx iris.Context) {
	s.Update <- true
	ctx.WriteString("Done")
}
