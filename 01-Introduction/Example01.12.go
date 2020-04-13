package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
func main() {
	// Type only
	var start, middle, end float32
	fmt.Println(start, middle, end)
	// Initial value mixed type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)
	// works with functions also
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
	/*
										Non-English Variable Names

		Go is a UTF-8 compliant language, which means you can define variables' names using alphabets
		other than the Latin alphabet that, for example, English uses
	*/
	fmt.Println("Using characters of non english alphabet")
	デバッグ := false
	日志级别 := "info"
	ይጀምሩ := time.Now()
	_A1_Μείγμα := ""
	fmt.Println(デバッグ, 日志级别, ይጀምሩ, _A1_Μείγμα)
}
