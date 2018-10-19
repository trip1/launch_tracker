package main

// Data types used when fetching
// info from launchlibrary.net
//
//

// Launch object for top level data
type Launch struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Start string `json:"isostart"`
}

// Launches array for loading
// response from launchlibrary
type Launches struct {
	Data []Launch `json:"launches"`
}
