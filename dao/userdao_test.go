package dao

import (
	"fmt"
	"testing"
)

func testCheckUserNameAndPassword(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("user = ", user)
}

func testCheckUserName(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("user =", user)
}

func testSaveUser(t *testing.T) {
	err := SaveUser("lvzhiqiang", "123456", "lvzhiqiang@grgbanking.com")
	if err != nil {
		err.Error()
	}
}
