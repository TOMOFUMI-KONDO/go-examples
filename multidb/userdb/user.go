package userdb

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func NewUser(id int, name string) *User {
	return &User{id, name}
}
