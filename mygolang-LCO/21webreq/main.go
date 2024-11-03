package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	fmt.Println("Web verb in go")
	PerformGetReq()

}

func PerformGetReq()  {
	const myurl = "http://127.0.0.1:4040/get"

	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status code", resp.StatusCode)
	fmt.Println("Content length", resp.ContentLength)

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}