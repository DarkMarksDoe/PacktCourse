package main

import (
	"fmt"
	"time"
)

func main() {
	/* VAR1, VAR2, VAR3 := VAL1, VAL2, VAL3
	 */
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	//This is called MULTI VARIABLE SHORT VARIABLE DECLARATION

	fmt.Println(Debug, LogLevel, startUpTime)
}
