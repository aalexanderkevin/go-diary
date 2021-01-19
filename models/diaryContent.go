package models

// Model Diary Content
type DiaryContent struct {
	Content string `json:"content" example:"This is content of the diary"`
	Date    string `json:"date" example:"2021-01-19"`
}
