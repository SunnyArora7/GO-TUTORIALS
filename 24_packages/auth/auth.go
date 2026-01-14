package auth

import "fmt"

// To access the method outside the folder package then the name of the function should start with capital letters
func Login(email string, password string) {
	fmt.Println("Logged in with email and Password:", email, password)
}
