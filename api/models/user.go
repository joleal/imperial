package models

//User represents a saved user
type User struct {
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}
