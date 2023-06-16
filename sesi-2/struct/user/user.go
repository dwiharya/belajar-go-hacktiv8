package user

import "fmt"

var Users []User

type User struct {
	Name     string
	Email    string
	Password string
}

func Register(newUser User) {
	Users = append(Users, newUser)
	fmt.Println("User registered successfully!")
}

func Get() []User {
	return Users
}
