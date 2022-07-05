package main

import (
	"fmt"
	"os"
	"strings"

	hw02unpackstring "github.com/aberestyak/golang-hw/hw02_unpack_string/pkg"
)

func main() {
	input := strings.Join(os.Args[1:], "")
	upackedString, err := hw02unpackstring.Unpack(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(upackedString)
}
