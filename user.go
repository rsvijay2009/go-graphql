package go_graphql

import (
	"time"
)

// User struct to define a user
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetAllUsers method to list all the users
func GetAllUsers() ([]*User, error) {
	var result []*User

	rows, err := DB.Query("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM `users`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u User

		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			return nil, err
		}

		result = append(result, &u)
	}

	return result, nil
}

// Create method to create a new user
func (u *User) Create() error {
	result, err := DB.Exec("INSERT INTO `users` (first_name, last_name, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", u.FirstName, u.LastName, u.Email, u.Password, time.Now(), time.Now())

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = lastID
	return nil
}
