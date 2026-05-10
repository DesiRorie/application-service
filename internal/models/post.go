package models

type Post struct {
	Company  string `json:"company"`
	Position string `json:"positon"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}
