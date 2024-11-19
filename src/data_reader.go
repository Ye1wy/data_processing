package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	FFlag = flag.String("f", "", "Read xml or json file")
)

type DBReader interface {
	ExtractData()
}

type XmlData struct {
	XMLName xml.Name `xml:"recipes"`
	cake    []Cake   `xml:"cake"`
}

type Cake struct {
	XMLName     xml.Name      `xml:"cake"`
	name        string        `xml:"name"`
	stovetime   string        `xml:"stovetime"`
	ingredients []Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	XMLName xml.Name `xml:"ingredients"`
	name    string   `xml:"itemname"`
	count   float32  `xml:"itemcount"`
	unit    string   `xml:"itemunit"`
}

type JsonData struct {
	name, stove_time string
	ingridients      []string
}

func (x XmlData) ExtractData() {
	XmlFile, err := os.Open(*FFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nFile opened\n")
	defer XmlFile.Close()

	byteValue, _ := io.ReadAll(XmlFile)
	err = xml.Unmarshal(byteValue, &x)

	if err != nil {
		fmt.Printf("[Error]: %v\n", err)
		return
	}

	fmt.Printf("%#v", x)
}

func main() {
	flag.Parse()

	fmt.Print("File: ", *FFlag)
	data := new(XmlData)
	data.ExtractData()
}
