package main

import "fmt"

func main()  {
	fmt.Println("Loops in golang")

	days := []string{"sun","mon","wed","thurs","fri"}
	fmt.Println(days)

	for i:=0; i<len(days); i++{
		fmt.Println(days[i])
	}
	
	for i:= range days{
		fmt.Println(days[i])
	}

	for index,days := range days{
		fmt.Println(index,days)
	}
	for _,days := range days{
		fmt.Println(days)
	}

	rougueValue := 1

	for rougueValue < 10{

		if rougueValue == 2{
			goto abc
		}

		if rougueValue == 5{
			rougueValue++
			continue
		}
		fmt.Println("value is", rougueValue)
		rougueValue++

		

	}
	abc: 
		fmt.Println("Jumping to abc")
			// rougueValue++
}