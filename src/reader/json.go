package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"./../data"
)

type JsonData struct {
	Cake []data.JsonCake `json:"cake"`
}

func (j *JsonData) Parse(file *os.File) error {
	byteValue, _ := io.ReadAll(file)
	err := json.Unmarshal(byteValue, &j)

	if err != nil {
		return err
	}

	bI, _ := json.MarshalIndent(j, "", " ")
	fmt.Println(string(bI))

	return nil
}

func (j *JsonData) ToCommon() CommonData {
	var common CommonData

	for _, c := range j.Cake {
		var ingredients []data.Ingredients

		for _, i := range c.Ingredients {
			ingredients = append(ingredients, data.Ingredients{
				Name:  i.Name,
				Count: i.Count,
				Unit:  i.Unit,
			})
		}

		common.data = append(common.data, data.Cake{
			Name:        c.Name,
			Time:        c.Time,
			Ingredients: ingredients,
		})
	}

	return common
}
