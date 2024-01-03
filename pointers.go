package main

import "fmt"

// pass by reference
func changeValue(a *int) {
	*a = 20
}


func main() {
	// pointer
	var a int = 10
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // dereferencing

	// pointer to pointer
	var c **int = &b
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(**c)

	// pass by reference
	changeValue(&a)
	fmt.Println(a)
}