package main

type Launch struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Start string `json:"windowstart"`
}

type Launches struct {
	Data []Launch `json:"launches"`
}
