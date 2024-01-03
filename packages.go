package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	greetings := "Hello there friends"

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)
	
}