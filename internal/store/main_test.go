package store

import (
	"github.com/gin-gonic/gin"
)

func testSetup() {
	gin.SetMode(gin.TestMode)
	ResetTestDatabase()
}

func addTestUser() (*User, error) {
	user := &User{
		Username: "batman",
		Password: "secret123",
	}
	err := AddUser(user)
	return user, err
}
