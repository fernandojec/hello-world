package main

import (
	"fmt"

	"github.com/fernandojec/hello-world/struct/user"
)

func main() {
	user.Register(user.User{
		Name:     "Andi",
		Email:    "andi@mail.com",
		Password: "123",
	})
	user.Register(user.User{
		Name:     "Andi1",
		Email:    "andi1@mail.com",
		Password: "1231",
	})
	user.Register(user.User{
		Name:     "Andi2",
		Email:    "andi2@mail.com",
		Password: "1232",
	})

	for _, v := range user.Get() {
		fmt.Printf("name: %s\nemail: %s\npassword: %s\n\n", v.Name, v.Email, v.Password)
	}

}
