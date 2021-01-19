package models

// Models API Response
type Response struct {
	Message string `json:"status" example:"Success"`
}

// Models API Error Response
type ErrorResponse struct {
	Message string `json:"error_message" example:"invalid JSON format"`
}

// Models JWT Response
type JWT struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTA4OTY5MjYsInVzZXJuYW1lIjoibGFsYSIsInV1aWQiOiI5YTA3YTIwYi01MTYwLTQ4N2ItYTBlYS1iMzBkZjM3NmMyMjcifQ.3kMJSHB-pMSjRovvVPU1O1p6Y04qgLDaJKr1ONPtkvY"`
}
