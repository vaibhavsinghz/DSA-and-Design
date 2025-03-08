package models

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(userID int, name, email string) *User {
	return &User{
		ID:    userID,
		Name:  name,
		Email: email,
	}
}
