package user

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User

func Register(user User) {
	users = append(users, user)
}

func Get() []User {
	return users
}
