package main

import "fmt"

func main()  {
	fmt.Println("welcome to learn pointer")
	
	// var ptr *int
	// fmt.Println("value of ptr", ptr)

	myNumber := 23
	var ptr = &myNumber //refrencing 
	fmt.Println("value of actual pointer", ptr) // actual memory address
	fmt.Println("value of actual pointer", *ptr) // value

	*ptr = *ptr + 2
	fmt.Println("New value is", myNumber)
	
}