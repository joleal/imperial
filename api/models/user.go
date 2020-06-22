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

//SearchUsers searches for users in the database filtering email and username by passed string
func SearchUsers(name string) (*[]*User, error) {
	db := CreateDb()
	sql := `

		SELECT 
			user_id AS id,
			nickname,
			email
		FROM user u
		WHERE 
			u.nickname LIKE ? 
			OR u.email LIKE ?
		LIMIT 10;`
	filter := "%" + name + "%"
	rows, err := db.Query(sql, filter, filter)
	if err != nil {
		log.Println(err)
	}

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Nickname, &user.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &users, nil
}
