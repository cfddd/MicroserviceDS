package model

import (
	"fmt"
	"testing"
)

func TestUserModel_Create(t *testing.T) {
	InitDb()
	user := &Users{
		Name:     "李四",
		Password: "123456",
	}
	GetInstance().Create(user)
}

func TestUserModel_FindUserByName(t *testing.T) {
	InitDb()
	user, _ := GetInstance().FindUserByName("张三")
	fmt.Print(user)
}
