package user

import "fmt"

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User

func Register(user User) {
	users = append(users, user)
}

func Get() {
	for _, v := range users {
		fmt.Printf("name: %s\nemail: %s\npassword: %s\n\n", v.Name, v.Email, v.Password)
	}
}
