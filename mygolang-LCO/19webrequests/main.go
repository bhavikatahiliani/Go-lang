package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"
func main()  {
	fmt.Println("Webrequest in Go")
	resp, err := http.Get(url)

	if err != nil{
		panic(err)
	}
	fmt.Printf("Resp is of type %T\n", resp)
	defer resp.Body.Close()

	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println("dataBytes", string(dataBytes))
}