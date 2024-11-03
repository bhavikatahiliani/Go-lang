package main

import (
	"fmt"
	"sort"
)


func main() {
	fmt.Println("welcome to slices")
	var fruitList = []string{"Apple", "Mango"}
	fmt.Printf("Type of fruitList %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Lichi")
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList[1:]) // removes first element
	fmt.Println("fruitList:", fruitList)
	fruitList = append(fruitList[1:3]) // start form 1 and ends at 2 as last range is non inclusive
	fmt.Println("fruitList:", fruitList)
	fruitList = append(fruitList[:1]) // starts from 0 
	fmt.Println("fruitList:", fruitList)

	highScores := make([]int,4)
	highScores[0] = 5
	highScores[1] = 4
	highScores[2] = 3
	highScores[3] = 2
	// highScores[4] = 1 // return error
	highScores = append(highScores, 1) // work dispite error above as it realocate the momory
	fmt.Println("highScores:", highScores)

	sort.Ints(highScores)
	fmt.Println("highScores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"js","java","c","python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //remove the given index
	fmt.Println(courses)



}