package main

import "fmt"

//Declaring a global string variable
var foo string = "bar"

func main() {
	//Declaring a local string variable
	var baz string = "qux"
	//Printing both variables
	fmt.Println(foo, baz)
}
