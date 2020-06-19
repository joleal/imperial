package models

import "log"

//User represents a saved user
type User struct {
	ID       string `json:"user_id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

//RegisterUser registers user in the database
func (u *User) RegisterUser() error {

	db := CreateDb()
	sql := `

		INSERT IGNORE INTO user 
		SET user_id = ?, 
			nickname = ?, 
			email =?;`

	_, err := db.Query(sql, u.ID, u.Nickname, u.Email)
	if err != nil {
		log.Println("couldn't decode body")
		return err
	}

	return nil
}
