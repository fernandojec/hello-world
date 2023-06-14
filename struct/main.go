package main

import "github.com/fernandojec/hello-world/struct/user"

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

	user.Get()
}
