package models

// Model Diary
type Diary struct {
	Content string `json:"content" example:"This is content of the diary"`
	Date    string `json:"date" example:"2021-01-12"`
}
