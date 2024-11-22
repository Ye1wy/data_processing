package main

import (
	"data_processing/src/compare"
	"data_processing/src/reader"
	"flag"
	"fmt"
	"os"
)

var (
	OldFlag = flag.String("old", "nothing", "First readed snapshot.txt file")
	NewFlag = flag.String("new", "nothing", "Second readed snapshot.txt file")
)

func main() {
	flag.Parse()

	if *OldFlag == "nothing" || *NewFlag == "nothing" {
		fmt.Println("[Error] File: Be better, one of file is nothing")
		return
	}

	old_file, err := os.Open(*OldFlag)

	if err != nil {
		fmt.Println("[Error] Coudn't open file")
		return
	}

	defer old_file.Close()

	new_file, err := os.Open(*NewFlag)

	if err != nil {
		fmt.Println("[Error] Coudn't open file")
		return
	}

	defer new_file.Close()

	data := reader.FSFileRead(old_file, new_file)

	fmt.Println(data.Old_file_data)
	fmt.Println(data.New_file_data)

	compare.FSCompare(data)
}
