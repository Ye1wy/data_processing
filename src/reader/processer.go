package reader

import (
	"io"
	"os"
)

func ProcessFile(file *os.File) (DBReader, error) {
	parser, err := DetectFileType(file)

	if err != nil {
		return nil, err
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return parser, err
}
