package main

import "fmt"

func main() { // Realization of mapes
	var dict_name_age = map[string]int{
		"Lev":     18,
		"Daniel":  22,
		"Michael": 36,
	}
	//Writing Daniels age
	fmt.Println(dict_name_age["Daniel"])
	dict_name_age["Daniel"] = 50 // Change age of Daniel
	fmt.Println(dict_name_age["Daniel"])
	delete(dict_name_age, "Daniel") // deleting pair - ("Daniel", 50)
	fmt.Println(dict_name_age["Daniel"])
	dict_name_age["Lev"]++ // increasing by 1 Lev's age
	fmt.Println(dict_name_age["Lev"])
}
