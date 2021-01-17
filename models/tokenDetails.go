package models

// Model Token Details
type TokenDetails struct {
	UUID     string
	Token    string
	Username string
	Exp      int64
}
