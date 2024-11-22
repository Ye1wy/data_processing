package reader

import (
	"data_processing/src/reader"
	"os"
	"testing"
)

func TestDetectingFile_0(t *testing.T) {
	file, _ := os.Open("files/json/stolen_database.json")

	out, err := reader.DetectFileType(file)

	var err_want error

	err_want = nil

	if err != err_want {
		t.Errorf("Func should not give an error")
	}

	if _, ok := out.(*reader.JsonData); !ok {
		t.Errorf("Func did not return the desired type")
	}
}

func TestDetectingFile_1(t *testing.T) {
	file, _ := os.Open("files/xml/original_database.xml")

	out, err := reader.DetectFileType(file)

	var err_want error

	err_want = nil

	if err != err_want {
		t.Errorf("Func should not give an error")
	}

	if _, ok := out.(*reader.XmlData); !ok {
		t.Errorf("Func did not return the desired type")
	}
}

func TestDetectingFile_2(t *testing.T) {
	file, _ := os.Open("files/snapshot/snapshot_1.txt")

	_, err := reader.DetectFileType(file)

	if err == nil {
		t.Errorf("Func should give error")
	}
}

func TestJson_0(t *testing.T) {
	file, _ := os.Open("files/json/stolen_database.json")

	var out reader.JsonData

	err := out.Parse(file)

	if err != nil {
		t.Errorf("Func should not give an error")
	}

	common := out.ToCommon()

	if common.Data[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Time != "45 min" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Name != "Flour" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Count != "2" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Unit != "mugs" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Name != "Strawberries" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Count != "8" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Unit != "" {
		t.Errorf("Data is corrupted")
	}
}

func TestJson_1(t *testing.T) {
	file, _ := os.Open("files/xml/original_database.xml")

	var out reader.JsonData

	err := out.Parse(file)

	if err == nil {
		t.Errorf("Func should give error")
	}
}

func TestXml_0(t *testing.T) {
	file, _ := os.Open("files/xml/original_database.xml")

	var out reader.XmlData

	err := out.Parse(file)

	if err != nil {
		t.Errorf("Func should not give an error")
	}

	common := out.ToCommon()

	if common.Data[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Time != "40 min" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Name != "Flour" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Count != "3" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[0].Unit != "cups" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Name != "Vanilla extract" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Count != "1.5" {
		t.Errorf("Data is corrupted")
	}

	if common.Data[0].Ingredients[1].Unit != "tablespoons" {
		t.Errorf("Data is corrupted")
	}
}

func TestXml_1(t *testing.T) {
	file, _ := os.Open("files/json/stolen_database.json")

	var out reader.XmlData

	err := out.Parse(file)

	if err == nil {
		t.Errorf("Func should give error")
	}
}

func TestReadLiens(t *testing.T) {
	file, _ := os.Open("files/snapshot/snapshot_1.txt")

	out, err := reader.ReadLines(file)

	if err != nil {
		t.Errorf("Func should not give an error")
	}

	want := []string{"/etc/stove/config.xml", "/Users/baker/recipes/database.xml", "/Users/baker/recipes/database_version3.yaml", "/var/log/orders.log", "/Users/baker/pokemon.avi"}

	for _, item := range want {
		if _, exist := out[item]; !exist {
			t.Errorf("Func should not give an error")
			break
		}
	}
}
