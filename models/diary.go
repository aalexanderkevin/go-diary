package models

// Model Diary
type Diary struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
