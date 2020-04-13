package main

import (
	"fmt"
	"time"
)

//Declaring a variable that have variables, like a MAP
var (
	//NAME     //TYPE	  //VALUE
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	//Printing the variables
	fmt.Println(Debug, LogLevel, startUpTime)
}
