package main

import "fmt"

func main(){
	// if else elseif
	x := 5
	if x > 5 {
		fmt.Println("x is greater than 5")
	}else if x < 5 {
		fmt.Println("x is less than 5")
	}else {
		fmt.Println("x is equal to 5")
	}
	
	names := []string{"sourav", "mario", "luigi", "peach", "bowser"}
	
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

	// switch case
	x = 3
	switch x {
	case 1:
		fmt.Println("x is equal to 1")
	case 2:
		fmt.Println("x is equal to 2")
	case 3:
		fmt.Println("x is equal to 3")
	default:
		fmt.Println("x is not equal to 1, 2 or 3")
	}
}