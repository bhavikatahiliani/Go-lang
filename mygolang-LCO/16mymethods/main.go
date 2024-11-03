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

	Bhavika.GetStatus()
	Bhavika.NewMail()
	fmt.Printf("name %v, email %v", Bhavika.Name, Bhavika.Email)
	Bhavika.NewMail2()
	fmt.Printf("name %v, email %v", Bhavika.Name, Bhavika.Email)
	
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User)  GetStatus(){
	fmt.Println("Is user Active?", u.Status)
	
}

func (u User)  NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("new email is", u.Email) // creates a copy and update
	
}
func (u *User)  NewMail2(){
	u.Email = "test@gmail.com"
	fmt.Println("new email is", u.Email) // updates actual value
	
}