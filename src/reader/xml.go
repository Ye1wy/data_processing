package reader

import (
	"encoding/xml"
	"io"
	"os"

	"data_processing/src/data"
)

type XmlData struct {
	XMLName xml.Name       `xml:"recipes"`
	Cake    []data.XmlCake `xml:"cake"`
}

func (x *XmlData) Parse(file *os.File) error {
	byteValue, _ := io.ReadAll(file)
	err := xml.Unmarshal(byteValue, &x)

	if err != nil {
		return err
	}

	return nil
}

func (x *XmlData) ToCommon() (CommonData, string) {
	var common CommonData

	for _, c := range x.Cake {
		var ingredients []data.Ingredients

		for _, i := range c.Ingredients.Item {
			ingredients = append(ingredients, data.Ingredients{
				Name:  i.Name,
				Count: i.Count,
				Unit:  i.Unit,
			})
		}

		common.Data = append(common.Data, data.Cake{
			Name:        c.Name,
			Time:        c.Stovetime,
			Ingredients: ingredients,
		})
	}

	return common, string("xml")
}
