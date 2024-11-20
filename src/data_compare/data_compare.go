package main

import (
	"data_processing/src/reader"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	OldFlag = flag.String("old", "nothing", "First readed file who will converted to new file")
	NewFlag = flag.String("new", "nothing", "Second readed file who will compared with new file")
)

func main() {
	flag.Parse()

	File, err := os.Open(*OldFlag)

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

	if _, err = File.Seek(0, io.SeekStart); err != nil {
		fmt.Printf("[Error] Resetting file: %v\n", err)
		return
	}

	out := parser.Parse(File)

}
