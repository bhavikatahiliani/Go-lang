package main
import (
	"fmt"
)

func main()  {
	fmt.Println("Structs:")
	Bhavika := User{"bhavi","bhavi@go.dev",true,23}
	fmt.Println(Bhavika)
	fmt.Printf("Bhavi's details are %+v\n", Bhavika)
	fmt.Printf("Name is %v and email is %v\n", Bhavika.Name, Bhavika.Email)
	
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}