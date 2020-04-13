package main

import "math/rand"

func main() {
	//create a seed, but we need a type int64 not a normal int
	var seed int64 = 1234456789
	//create seed with value int64
	rand.Seed(seed)
}
