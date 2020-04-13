package main

import (
	"bufio"
	"fmt"
	"os"
)

/*				IN THIS CODE WE SIMULATE A RESTAURANT			*/
func main() {
	//Main course, Total = 2
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)
	// Drinks Total = 4
	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)
	// Discount
	total = total - 5
	fmt.Println("Sub :", total)
	// 10% Tip
	tip := total * 0.10
	fmt.Println("Tip :", tip)
	//Total to pay
	total = total + tip
	fmt.Println("Total:", total)
	// Split bill, there was 2 persons, so...
	split := total / 2
	fmt.Println("Split:", split)
	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1
	//So visitCount is now 25
	/* REMAINDER checks every 5 persons with MOD 5*/
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward babe!!. Tell us your name: ")
		/*CREATING AN INPUT READING FROM USER*/
		//Declare de Standard Input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()         // use `for scanner.Scan()` to keep reading
		name := scanner.Text() //converts to text
		fmt.Println("Congratulations", name, "you won a 25%", "for a new visit!")
	}

}
