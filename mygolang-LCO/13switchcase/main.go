package main

import (
	"math/rand"
	"time"
	"fmt"
)


func main()  {
	fmt.Println("Switch Case in golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice is :" , diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	case 4:
		fmt.Println("Value is 4")
		fallthrough
	case 5:
		fmt.Println("Value is 5")
		fallthrough
	case 6:
		fmt.Println("Value is 6")
	default:
		fmt.Println("Unknown value")
	}
	
}