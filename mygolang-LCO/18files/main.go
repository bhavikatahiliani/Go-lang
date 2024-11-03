package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Files in Golang")
	content := "This needs to go inside the file"
	file, err := os.Create("./mygofile.txt")
	if err != nil{
		panic(err)
	}
	length, err := io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filname string)  {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)
	fmt.Println("data inside file", string(databyte))
	
}

func checkNilErr(err error)  {
	if err != nil{
		panic(err)
	}
}