package main

import (
	"fmt"

	"github.com/koskuma/gotools"
	"github.com/koskuma/gotools/stringhelper"
)

func init() {
	fmt.Println("Init Ja!")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("cat"))
}
