package main

import (
	"fmt"
	"stringmodify"
)

func main() {
	helloworld := "Hello"
	stringmodify.Stringworld(&helloworld)
	fmt.Println(helloworld)

}
