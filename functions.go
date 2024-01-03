package main

import "fmt"

func sayGreetings(name string){
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string){
	fmt.Printf("Good bye %v \n", name)
}

func add(x int, y int) int{
	return x + y
}

func subtract(x int, y int) int{ // or func subtract(x, y int) int = return type is int
	return x - y
}

func addSubtract(x int, y int) (int, int){
	return x + y, x - y
}


// array as parameer
func addSubtract(x [2]int, y [2]int) ([2]int, [2]int){
	var sum [2]int
	var diff [2]int
	sum[0] = x[0] + y[0]
	sum[1] = x[1] + y[1]
	diff[0] = x[0] - y[0]
	diff[1] = x[1] - y[1]
	return sum, diff
}



func main(){

	sayGreetings("Sourav")
	sayGreetings("Mario")
	sayGreetings("Luigi")
	sayBye("Sourav")
	sayBye("Mario")
	sayBye("Luigi")

	var x int = 5
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)

	var x int = 5
	var y int = 7
	var sum int = add(x, y)
	fmt.Println(sum)

	var x int = 5
	var y int = 7
	var diff int = subtract(x, y)
	fmt.Println(diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)

	var x int = 5
	var y int = 7
	var sum, diff int = addSubtract(x, y)
	fmt.Println(sum, diff)
}