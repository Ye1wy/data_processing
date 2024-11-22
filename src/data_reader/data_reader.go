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
	FFlag = flag.String("f", "nothing", "Read xml or json file")
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

	if *FFlag == "nothing" {
		fmt.Println("Did nothing! Be better my dear l*ser... *kxm* User!")
		return
	}

	File, err := os.Open(*FFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer File.Close()

	parser, err := reader.DetectFileType(File)

	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}

	if _, err = File.Seek(0, io.SeekStart); err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}

	err = parser.Parse(File)

	if err != nil {
		fmt.Printf("[Error] Parsing file: %v\n", err)
		return
	}

	out := parser.ToCommon()

	if _, ok := parser.(*reader.XmlData); ok {
		PrettyPrintJson(out)

	} else {
		PrettyPrintXml(out)
	}
}
