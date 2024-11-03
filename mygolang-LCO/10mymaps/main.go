package main

import (
	"fmt"
)

func main()  {

	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["Js"] = "javasript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("list of all languages", languages)
	fmt.Println("Js", languages["js"])
	delete(languages,"rb")
	fmt.Println("list of all languages", languages)

	for key, value := range languages{
		fmt.Printf("for key %v, value as %v\n", key, value)
	}

	for _, value := range languages{
		fmt.Printf(" value as %v\n",  value)
	}


	
}