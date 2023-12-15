package userdb

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	return &User{id, name}
}
