package main

import "fmt"

func main(){

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40


	// bits	& memory
	var numOne int16 = 255

	// format specifier
	fmt.Printf("The age is %d\n", ageOne)
	fmt.Printf("The age is %d\n", ageTwo)
	fmt.Printf("The age is %d\n", ageThree)
	fmt.Print("This is a string\n")	
	fmt.Printf("Formate specifiers are %T %v", numOne, numOne)
}
