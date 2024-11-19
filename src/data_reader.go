package main

import (
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

type FileParser interface {
	Parse(file *os.File) error
}

type XmlData struct {
	XMLName xml.Name `xml:"recipes"`
	Cake    []Cake   `xml:"cake"`
}

type Cake struct {
	XMLName     xml.Name      `xml:"cake"`
	Name        string        `xml:"name"`
	Stovetime   string        `xml:"stovetime"`
	Ingredients []Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	XMLName xml.Name `xml:"ingredients"`
	Name    string   `xml:"itemname"`
	Count   float32  `xml:"itemcount"`
	Unit    string   `xml:"itemunit"`
}

type JsonData struct {
	name, stove_time string
	ingridients      []string
}

func (x *XmlData) Parse(file *os.File) error {

	byteValue, _ := io.ReadAll(file)
	err := xml.Unmarshal(byteValue, &x)

	if err != nil {
		return err
	}

	fmt.Printf("%#v", x)
	return nil
}

func (j *JsonData) Parse(file *os.File) error {
	return nil
}

func DetectFileType(file *os.File) (FileParser, error) {
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
