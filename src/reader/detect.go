package reader

import (
	"fmt"
	"os"
	"strings"
)

func DetectFileType(file *os.File) (DBReader, error) {
	peek := make([]byte, 512)
	_, err := file.Read(peek)

	if err != nil {
		return nil, fmt.Errorf("[Error] Error reading file: %w\n", err)
	}

	if strings.HasPrefix(strings.TrimSpace(string(peek)), "{") {
		return &JsonData{}, nil
	}

	if strings.HasPrefix(strings.TrimSpace(string(peek)), "<") {
		return &XmlData{}, nil
	}

	return nil, fmt.Errorf("Unknown file type")
}
