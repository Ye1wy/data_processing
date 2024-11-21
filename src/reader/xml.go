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

func (xml_data *XmlData) Parse(file *os.File) error {
	byteValue, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = xml.Unmarshal(byteValue, &xml_data)

	if err != nil {
		return err
	}

	return nil
}

func (xml_data *XmlData) ToCommon() *CommonData {
	common := &CommonData{}

	for _, cake := range xml_data.Cake {
		var ingredients []data.Ingredients

		for _, ingredients_item := range cake.Ingredients.Item {
			ingredients = append(ingredients, data.Ingredients{
				Name:  ingredients_item.Name,
				Count: ingredients_item.Count,
				Unit:  ingredients_item.Unit,
			})
		}

		common.Data = append(common.Data, data.Cake{
			Name:        cake.Name,
			Time:        cake.Stovetime,
			Ingredients: ingredients,
		})
	}

	return common
}
