package data

type JsonCake struct {
	Name        string            `json:"name"`
	Time        string            `json:"time"`
	Ingredients []JsonIngredients `json:"ingredients"`
}

type JsonIngredients struct {
	Name  string `json:"ingredient_name"`
	Count string `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit"`
}
