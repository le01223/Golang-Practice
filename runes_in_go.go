package main

import (
	"fmt"
	"reflect"
)

func main() { //Working with runes
	message := []rune("I am a good programmer😎")
	for _, str := range message {
		fmt.Println(str, string(str)) // write code symbol and string performance
	}
	rune2 := '🦍'
	fmt.Println(rune2, reflect.TypeOf(rune2)) // OMG type is int32
}
