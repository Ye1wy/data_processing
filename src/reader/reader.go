package reader

import (
	"fmt"
	"os"

	"data_processing/src/data"
)

type CommonData struct {
	Data []data.Cake `json:"cake" xml:"cake"`
}

type DBReader interface {
	Parse(file *os.File) (*CommonData, error)
}

func FileReader(data DBReader, file *os.File) *CommonData {
	out, err := data.Parse(file)

	if err != nil {
		fmt.Printf("[Error] Parsing file: %v\n", err)
		return nil
	}

	return out
}
