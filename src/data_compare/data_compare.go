package main

import (
	"data_processing/src/compare"
	"data_processing/src/reader"
	"flag"
	"fmt"
	"os"
)

var (
	OldFlag = flag.String("old", "nothing", "First readed file who will converted to new file")
	NewFlag = flag.String("new", "nothing", "Second readed file who will compared with new file")
)

func main() {
	flag.Parse()

	if *OldFlag == "nothing" || *NewFlag == "nothing" {
		fmt.Println("[Error] Both --old and --new flags are required")
		return
	}

	old_file, err := os.Open(*OldFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer old_file.Close()

	parser, err := reader.ProcessFile(old_file)

	if err != nil {
		fmt.Printf("[Error] %v", err)
		return
	}

	err = parser.Parse(old_file)

	if err != nil {
		fmt.Printf("[Error] Parsing: %v\n", err)
		return
	}

	old_data := parser.ToCommon()

	new_file, err := os.Open(*NewFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer new_file.Close()

	parser, err = reader.ProcessFile(new_file)

	if err != nil {
		fmt.Printf("[Error] %v", err)
		return
	}

	err = parser.Parse(new_file)

	if err != nil {
		fmt.Printf("[Error] Parsing: %v\n", err)
		return
	}

	new_data := parser.ToCommon()

	compare.DataCompare(old_data, new_data)
}
