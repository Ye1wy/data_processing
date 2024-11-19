package data

import "encoding/xml"

type XmlCake struct {
	XMLName     xml.Name         `xml:"cake"`
	Name        string           `xml:"name"`
	Stovetime   string           `xml:"stovetime"`
	Ingredients []XmlIngredients `xml:"ingredients"`
}

type XmlIngredients struct {
	XMLName xml.Name `xml:"ingredients"`
	Name    string   `xml:"itemname"`
	Count   string   `xml:"itemcount"`
	Unit    string   `xml:"itemunit"`
}
