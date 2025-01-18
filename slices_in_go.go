package main

import "fmt"

func main() { // Creating slices
	languages := []string{"English", "Russian", "German"}
	fmt.Println(languages)
	nil_slice := []int{}
	fmt.Println(nil_slice)
	var bad_slice []int
	fmt.Println(bad_slice)
	make_slice := make([]int, 5, 20)   // (type, lenght, capacity)
	make_slice = append(make_slice, 1) // lenght = 6 cap = 20
	fmt.Println(make_slice)
	make_slice[0] = 10
	fmt.Println(make_slice) // first element got 10\
}
