package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("welcome to time study")
	presentTime :=time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1999, time.April, 8,23,23,0,0,time.UTC)
	fmt.Println("my birthday", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}