package main

import (
	"loginPackage/auth"
	"loginPackage/user"

	"github.com/fatih/color"
)

func main() {
	auth.Login("sunny@gmail.com", "sunny")
	newUser := user.User{Email: "sunny@gmai.com", Password: "12344"}
	color.Cyan(newUser.Email)
}
