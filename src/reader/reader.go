package reader

import (
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
