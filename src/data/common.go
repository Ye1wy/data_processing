package data

type Cake struct {
	Name        string
	Time        string
	Ingredients []Ingredients
}

type Ingredients struct {
	Name  string
	Count string
	Unit  string
}
