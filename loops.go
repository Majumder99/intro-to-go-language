package main

import "fmt"

func main(){
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	x := [4]string{"a", "b", "c", "d"}
	for i, v := range x {
		fmt.Println(i, v)
	}

	// key, value loop
	y := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range y {
		fmt.Println(k, v)
	}
	

}