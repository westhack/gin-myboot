package model

type Module struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
