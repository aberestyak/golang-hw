package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const helloWorldMessage = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(helloWorldMessage))
}
