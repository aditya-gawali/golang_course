package main

import (
	"fmt"

	"github.com/aditya-gawali/podcast/auth"
	"github.com/aditya-gawali/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("aditya123", "12345")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Name:  "aditya",
		Email: "aditya@gmail.com",
	}

	fmt.Println(user)
	color.Red(user.Email)

}
