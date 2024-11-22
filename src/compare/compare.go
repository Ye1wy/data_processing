package compare

import (
	"data_processing/src/data"
	"data_processing/src/reader"
	"fmt"
	"slices"
	"strings"
)

func DataCompare(old_data, new_data *reader.CommonData) {
	old_cakes := make(map[string]data.Cake)
	new_cakes := make(map[string]data.Cake)

	for _, cake := range old_data.Data {
		old_cakes[cake.Name] = cake
	}

	for _, cake := range new_data.Data {
		new_cakes[cake.Name] = cake
	}

	for name, _ := range new_cakes {
		if _, exist := old_cakes[name]; !exist {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}

	for name, _ := range old_cakes {
		if _, exist := new_cakes[name]; !exist {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}

	for name, old_cake := range old_cakes {
		if new_cake, exist := new_cakes[name]; exist {

			if old_cake.Time != new_cake.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, old_cake.Time, new_cake.Time)
			}

			old_ingredients := make(map[string]data.Ingredients)
			new_ingredients := make(map[string]data.Ingredients)

			for _, ing := range old_cake.Ingredients {
				old_ingredients[ing.Name] = ing
			}

			for _, ing := range new_cake.Ingredients {
				new_ingredients[ing.Name] = ing
			}

			for ing_name, _ := range new_ingredients {
				if _, exist := old_ingredients[ing_name]; !exist {
					fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ing_name, name)
				}
			}

			for ing_name, _ := range old_ingredients {
				if _, exist := new_ingredients[ing_name]; !exist {
					fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ing_name, name)
				}
			}

			for ing_name, old_ing := range old_ingredients {
				if new_ing, exists := new_ingredients[ing_name]; exists {
					if old_ing.Count != new_ing.Count {
						fmt.Printf("CHANGED uint count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
							ing_name, name, new_ing.Count, old_ing.Count)
					}

					if old_ing.Unit != "" && new_ing.Unit == "" {
						fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
							old_ing.Unit, ing_name, name)

					} else if old_ing.Unit != new_ing.Unit {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
							ing_name, name, new_ing.Unit, old_ing.Unit)
					}

				}
			}
		}
	}
}

func FSCompare(data data.FSData) {
	old_substrings := strings.Split(data.Old_file_data, "\n")
	new_substrings := strings.Split(data.New_file_data, "\n")

	for i := 0; i < len(old_substrings); i++ {
		contained := slices.Contains(new_substrings, old_substrings[i])

		if !contained {
			fmt.Printf("REMOVED %s\n", old_substrings[i])
		}
	}

	for i := 0; i < len(new_substrings); i++ {
		contained := slices.Contains(old_substrings, new_substrings[i])

		if !contained {
			fmt.Printf("ADDED %s\n", new_substrings[i])
		}
	}

}
