package data

type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"time"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients"`
}

type Ingredients struct {
	Name  string `json:"ingredient_name" xml:"ingredient_name"`
	Count string `json:"ingredient_count" xml:"ingredient_count"`
	Unit  string `json:"ingredient_unit" xml:"ingredient_unit"`
}
