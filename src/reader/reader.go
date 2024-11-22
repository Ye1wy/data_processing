package reader

import (
	"bufio"
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

func ReadLines(file *os.File) (map[string]struct{}, error) {
	lines := make(map[string]struct{})
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines[scanner.Text()] = struct{}{}
	}

	return lines, scanner.Err()
}
