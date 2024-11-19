package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"./reader"
)

var (
	FFlag = flag.String("f", "", "Read xml or json file")
)

func main() {
	flag.Parse()

	File, err := os.Open(*FFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nFile opened\n")
	defer File.Close()

	parser, err := reader.DetectFileType(File)

	if err != nil {
		fmt.Printf("[Error] Detecting file: %v\n", err)
		return
	}

	_, err = File.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Printf("[Error] Resetting file: %v\n", err)
		return
	}

	err = parser.Parse(File)

	if err != nil {
		fmt.Printf("[Error] Parsing file: %v\n", err)
		return
	}
}
