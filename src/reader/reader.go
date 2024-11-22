package reader

import (
	"fmt"
	"io"
	"os"

	"data_processing/src/data"
)

type CommonData struct {
	Data []data.Cake `json:"cake" xml:"cake"`
}

type DBReader interface {
	Parse(file *os.File) error
	ToCommon() *CommonData
}

func FSFileRead(old_file, new_file *os.File) data.FSData {
	var data data.FSData
	buf := make([]byte, 32*1024)

	for {
		byteValue, err := old_file.Read(buf)

		if byteValue > 0 {
			data.Old_file_data += string(buf[:byteValue])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("[Error]: Reading error: %v\n", err)
			break
		}
	}

	for {
		byteValue, err := new_file.Read(buf)

		if byteValue > 0 {
			data.New_file_data += string(buf[:byteValue])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("[Error]: Reading error: %v\n", err)
			break
		}
	}

	return data
}
