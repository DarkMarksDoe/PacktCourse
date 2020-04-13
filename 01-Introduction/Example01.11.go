package main

import (
	"fmt"
	"time"
)

//NameFunction   Type1, Type2 , Type3
func getConfig() (bool, string, time.Time) {
	//return   //VAL1, VAL2, VAL3
	return false, "info", time.Now()
}

func main() {
	//Var1,  Var2    , Var2    := call of the function
	Debug, LogLevel, startUpTime := getConfig()
	//Print the variables
	fmt.Println(Debug, LogLevel, startUpTime)
}
