package store

type User struct {
	Username string `binding:"required,min=5,max=30"`
	Password string `binding:"required,min=7,max=32"`
}

var Users []*User
