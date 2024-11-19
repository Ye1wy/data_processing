package reader

import (
	"os"

	"./../data"
)

type CommonData struct {
	data []data.Cake
}

type DBReader interface {
	Parse(file *os.File) error
	ToCommon() CommonData
}
