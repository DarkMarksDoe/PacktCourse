package main

import (
	"fmt"
	"time"
)

//Declaring a variable that have variables, like a MAP
var (
	//name		//Type
	Debug bool
	//name		//Value
	LogLevel = "info"
	//name		//Value
	startUpTime = time.Now()

	//Doesn't matter if you assign the type of varibale or not
)

func main() {
	//Printing the variables
	fmt.Println(Debug, LogLevel, startUpTime)
}
