package main

import (
	"flag"
	"fmt"
)

var (
	FFlag = flag.String("f", "", "Read xml or json file")
)

type DBReader interface {
	ExtractData()
}

func main() {
	flag.Parse()

	fmt.Print("File: ", *FFlag)
}
