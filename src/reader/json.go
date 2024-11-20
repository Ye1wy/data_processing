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

func (j *JsonData) Parse(file *os.File) error {
	byteValue, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonData) ToCommon() *CommonData {
	common := &CommonData{}

	for _, c := range j.Cake {
		var ingredients []data.Ingredients

		for _, i := range c.Ingredients {
			ingredients = append(ingredients, data.Ingredients{
				Name:  i.Name,
				Count: i.Count,
				Unit:  i.Unit,
			})
		}

		common.Data = append(common.Data, data.Cake{
			Name:        c.Name,
			Time:        c.Time,
			Ingredients: ingredients,
		})
	}

	return common
}
