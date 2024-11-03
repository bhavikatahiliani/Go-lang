package main

import "fmt"

func main()  {

	fmt.Println("Functions in golang")
	greeter()
	result:= adder(3,5)

	fmt.Println("result", result)

	proRes, myMessage:= proAdder(2,5,8,7,3)
	fmt.Println("proRes", proRes)
	fmt.Println("proRes", myMessage)
	
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
	
}

func proAdder(values ...int) (int,string) {
	total := 0

	for _, val := range values{
		total+=val
	}
	return total, "this is my message"
}

func greeter()  {
	fmt.Println("Namaste")
	
}

func greeterTwo()  {
	fmt.Println("Another func")
	
}
// func (){
// 	fmt.Println("Anonymus func")
	
// }()