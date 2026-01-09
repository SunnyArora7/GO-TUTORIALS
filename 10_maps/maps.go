package main

import "fmt"

func main() {
	//syntax
	m := make(map[string]string)

	fmt.Println(m)

	//To set an element in map
	m["name"] = "sunny"

	//To get an element from the maps
	fmt.Println(m["name"])

	//If we try to get non existing key element then it will be "" in case of string , 0 int , false, bool
	fmt.Println(m["sunny"])
	fmt.Println(len(m))

	//delete(m, "name")
	fmt.Println(m)

	//To check whether an element exist or not
	v, ok := m["name"]
	fmt.Println(v)  //it will return the value of the key
	fmt.Println(ok) //it will return true or false whether the element exist or not

	if ok {
		fmt.Println("Ok")
	} else {
		fmt.Println("Not Ok")
	}

}
