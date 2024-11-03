package main

import "fmt"

func main()  {
	fmt.Println("welcome to aaray in go lang")
	var fruitList [4]string
	fruitList[0] = "blueberry"
	fruitList[1] = "lichi"
	fruitList[3] = "peach"

	fmt.Println("fruit list", fruitList)
	fmt.Println("fruit list len", len(fruitList))

	var vegList = [5]string{"brocoli", "beans", "mashroom"}
	fmt.Println("vegiiee liest", vegList)
	fmt.Println("vegiiee liest", len(vegList))
}