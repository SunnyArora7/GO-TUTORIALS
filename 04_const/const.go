package main

import "fmt"

// we can declar const variable here also
const versionNumber int = 20

// const ver:=20 this one not allowed outside the method
func main() {

	//syntax
	const name string = "goLang"
	fmt.Println(name)

	// we can declar a group of const variables like this
	const (
		portNumber = 20
		host       = "localHost"
	)

	fmt.Println(portNumber, host)

}
