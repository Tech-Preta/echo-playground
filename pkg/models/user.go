package models

import "time"

// User representa um usuário no sistema
type User struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Email   string `json:"email" xml:"email"`
	Age     int    `json:"age" xml:"age"`
	Created string `json:"created" xml:"created"`
}

// NewUser cria um novo usuário com timestamp
func NewUser(name, email string, age int) *User {
	return &User{
		Name:    name,
		Email:   email,
		Age:     age,
		Created: time.Now().Format(time.RFC3339),
	}
}

// SetID define o ID do usuário
func (u *User) SetID(id int) {
	u.ID = id
}
