package main
import "fmt"

func main()  {

	fmt.Println("If else statements")
	loginCount := 23

	var result string
	
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login counts"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("more than 10")
	}
}