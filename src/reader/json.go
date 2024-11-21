package reader

import (
	"encoding/json"
	"io"
	"os"

	"data_processing/src/data"
)

type JsonData struct {
	Cake []data.JsonCake `json:"cake"`
}

func (json_data *JsonData) Parse(file *os.File) error {
	byteValue, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &json_data)

	if err != nil {
		return err
	}

	return nil
}

func (json_data *JsonData) ToCommon() *CommonData {
	common := &CommonData{}

	for _, cake := range json_data.Cake {
		var ingredients []data.Ingredients

		for _, ingredients_item := range cake.Ingredients {
			ingredients = append(ingredients, data.Ingredients{
				Name:  ingredients_item.Name,
				Count: ingredients_item.Count,
				Unit:  ingredients_item.Unit,
			})
		}

		common.Data = append(common.Data, data.Cake{
			Name:        cake.Name,
			Time:        cake.Time,
			Ingredients: ingredients,
		})
	}

	return common
}
