package models

// Models API Response
type Response struct {
	Message string `json:"error_message"`
}

// Models JWT Response
type JWT struct {
	Token string `json:"token"`
}
