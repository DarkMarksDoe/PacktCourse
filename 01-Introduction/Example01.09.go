package main

import (
	"fmt"
	"time"
)

func main() {
	//Short variable declaration
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	//printing the variables
	fmt.Println(Debug, LogLevel, startUpTime)
}

/*
The := shorthand is very popular with Go developers...
and the most common way in which variables get defined in real-world Go code
*/
