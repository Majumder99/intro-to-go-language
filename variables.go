package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	// strings
	var nameOne string = "mario" //explicitly defining the type
	var nameTwo = "luigi" //implicitly defining the type
	nameFour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameFour)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits	& memory
	var numOne int16 = 255
	var num = 12.23
	numThree := 1.5

	fmt.Println(numOne, num, numThree)
}
