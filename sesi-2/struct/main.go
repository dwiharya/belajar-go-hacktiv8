package main

import (
	"fmt"
	"golang/sesi-2/struct/user"
	// "github.com/dwiharya/belajar-go-hacktiv8/tree/main/sesi-2/struct/user"
)

func main() {
	user1 := user.User{Name: "Dwi", Email: "dwih@example.com", Password: "password1"}
	user2 := user.User{Name: "Jane Smith", Email: "jane@example.com", Password: "password2"}
	user3 := user.User{Name: "Bob Johnson", Email: "bob@example.com", Password: "password3"}

	user.Register(user1)
	user.Register(user2)
	user.Register(user3)

	users := user.Get()
	for _, u := range users {
		fmt.Printf("Name: %s, Email: %s, Password: %s\n", u.Name, u.Email, u.Password)
	}
}
