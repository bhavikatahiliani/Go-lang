package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating")

	// comma ok // err ok

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating" , input)
	fmt.Printf("Type %T " , input)

	fmt.Println("enter age")

	age , err := reader.ReadString('\n')
	fmt.Println("your age" , age)
	fmt.Println("err" , err)
	fmt.Printf("Type %T " , age)
}