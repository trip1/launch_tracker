package main

// Data types used when fetching
// info from launchlibrary.net
//
//

// Launch object for top level data
type Launch struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Start  string `json:"windowstart"`
	Rocket Rocket `json:"rocket"`
}

// Launches array for loading
// response from launchlibrary
type Launches struct {
	Data []Launch `json:"launches"`
}

// Rocket data type
type Rocket struct {
	Name   string `json:"name"`
	Config string `json:"configuration"`
	Wiki   string `json:"wikiurl"`
	Image  string `json:"imageURL"`
}
