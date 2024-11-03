package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjsjkskj"

func main()  {
	fmt.Println("Handling Urls in Golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("type of query params %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _,value := range qparams{
		fmt.Println("Params is: ", value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=Bhavi",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println("anotherUrl",anotherUrl)
}