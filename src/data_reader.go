package main

import (
	"data_processing/src/reader"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	FFlag = flag.String("f", "", "Read xml or json file")
)

func PrettyPrintJson(data interface{}) {
	json_data, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(json_data))
}

func PrettyPrintXml(data interface{}) {
	xml_data, _ := xml.MarshalIndent(data, "", "    ")
	fmt.Println(string(xml_data))
}

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

	if _, err = File.Seek(0, io.SeekStart); err != nil {
		fmt.Printf("[Error] Resetting file: %v\n", err)
		return
	}

	out := reader.FileReader(parser, File)

	if _, ok := parser.(*reader.XmlData); ok {
		PrettyPrintXml(out)

	} else {
		PrettyPrintJson(out)
	}
}
