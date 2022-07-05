package main

import (
	"fmt"
	"github.com/aberestyak/golang-hw/hw02_unpack_string/pkg"
	"os"
	"strings"
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
