package main

import (
	"fmt"
	"math"
)

func main() { //Несколько вариантов объявления массивов в Go
	// Массив из пяти 0 по умолчанию
	var nonInitializedArray [5]int
	for _, i := range nonInitializedArray {
		fmt.Println(i)
	}
	// инициализированный массив с указанной длиной
	var InitializedArrayWithLen [3]string = [3]string{"Moscow", "Paris", "New York"}
	for _, j := range InitializedArrayWithLen {
		fmt.Println(j)
	}
	// инициализированный массив с без указания длины
	InitializedArrayWithOutLen := [...]float64{1.2, 2.22222, math.Pi}
	for _, k := range InitializedArrayWithOutLen {
		fmt.Println(k)
	}
}
