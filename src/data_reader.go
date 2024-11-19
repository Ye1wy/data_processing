package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	FFlag = flag.String("f", "", "Read xml or json file")
)

type DBReader interface {
	Parse(file *os.File) error
}

type XmlData struct {
	XMLName xml.Name  `xml:"recipes"`
	Cake    []XmlCake `xml:"cake"`
}

type XmlCake struct {
	XMLName     xml.Name         `xml:"cake"`
	Name        string           `xml:"name"`
	Stovetime   string           `xml:"stovetime"`
	Ingredients []XmlIngredients `xml:"ingredients"`
}

type XmlIngredients struct {
	XMLName xml.Name `xml:"ingredients"`
	Name    string   `xml:"itemname"`
	Count   string   `xml:"itemcount"`
	Unit    string   `xml:"itemunit"`
}

type JsonData struct {
	Cake []JsonCake `json:"cake"`
}

type JsonCake struct {
	Name        string            `json:"name"`
	Time        string            `json:"time"`
	Ingredients []JsonIngredients `json:"ingredients"`
}

type JsonIngredients struct {
	Name  string `json:"ingredient_name"`
	Count string `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit"`
}

type CommonData struct {
	data []Cake
}

type Cake struct {
	name        string
	time        string
	ingredients []Ingredients
}

type Ingredients struct {
	name  string
	count string
	unit  string
}

func (x *XmlData) Parse(file *os.File) error {
	byteValue, _ := io.ReadAll(file)
	err := xml.Unmarshal(byteValue, &x)

	if err != nil {
		return err
	}

	bI, _ := xml.MarshalIndent(x, "", " ")
	fmt.Println(string(bI))

	return nil
}

func (j *JsonData) Parse(file *os.File) error {
	byteValue, _ := io.ReadAll(file)
	err := json.Unmarshal(byteValue, &j)

	if err != nil {
		return err
	}

	bI, _ := json.MarshalIndent(j, "", " ")
	fmt.Println(string(bI))

	return nil
}

func DetectFileType(file *os.File) (DBReader, error) {
	reader := io.NewSectionReader(file, 0, 512)
	peek := make([]byte, 512)
	_, err := reader.Read(peek)

	if err != nil && err != io.EOF {
		return nil, err
	}

	content := strings.TrimSpace(string(peek))

	if strings.HasPrefix(content, "{") || strings.HasPrefix(content, "[") {
		return &JsonData{}, nil

	} else if strings.HasPrefix(content, "<") || strings.HasPrefix(content, "<?xml") {
		return &XmlData{}, nil
	}

	return nil, fmt.Errorf("Unknown file type")
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

	parser, err := DetectFileType(File)

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
