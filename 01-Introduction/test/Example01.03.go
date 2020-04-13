package main

// Import extra functionality from packages

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//Create a random seed using the time
	rand.Seed(time.Now().UnixNano())
	//Create a random number between 0 and n
	rndNmbr := rand.Intn(5) + 1
	//Repeat string from the variable
	stars := strings.Repeat("*", rndNmbr)
	//Print the last String variable
	fmt.Println(stars)
}
