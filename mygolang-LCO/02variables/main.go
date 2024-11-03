package main

import "fmt"
const Token string = "hdjandjk" //public
func main()  {
	var username string = "Bhavika"
	fmt.Println(username)
	fmt.Printf("var type is:  %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("var type is: %T \n", isLoggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("var type is: %T \n", smallval)

	var smallfloat float32 = 255.4555
	fmt.Println(smallfloat)
	fmt.Printf("var type is: %T \n", smallfloat)

	var anotherVar int 
	fmt.Println(anotherVar)
	fmt.Printf("var type is: %T \n", anotherVar)

	var str string 
	fmt.Println(str)
	fmt.Printf("var type is: %T \n", str)

	var name = "bhavika"
	fmt.Println(name)

	number := 20
	fmt.Println(number)

	fmt.Println(Token)
	fmt.Printf("var type is: %T \n", Token)
}