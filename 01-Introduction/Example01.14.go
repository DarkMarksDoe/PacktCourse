package main

import "fmt"

func main() {
	//var1, var2, var3 := val1, val2, val3
	query, limit, offset := "bat", 10, 0
	//var2, var3, var3 := newValue1, var3, newValue2
	query, limit, offset = "ball", offset, 20
	//print de values
	fmt.Println(query, limit, offset)
}
