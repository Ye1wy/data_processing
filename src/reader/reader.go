package reader

import (
	"os"

	"data_processing/src/data"
)

type CommonData struct {
	data []data.Cake
}

type DBReader interface {
	Parse(file *os.File) error
	ToCommon() CommonData
}
