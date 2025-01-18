package main

import "fmt"

type person struct {
	name   string
	age    int
	salary int
}

func main() {
	emp := person{}      // Empty person
	fmt.Println(emp.age) //Default age = 0
	emp2 := person{name: "Daniel", age: 18, salary: 20000}
	fmt.Println(emp2.salary)
	//Using Pointers
	alex := person{name: "Alex", age: 20}
	var alexPointer *person = &alex
	alexPointer.age -= 10 //alex.age -= 10 too
	fmt.Println(alex.age)
	//One more example
	lev := person{name: "Lev", age: 52}
	var agePointer *int = &lev.age
	*agePointer += 20 // lev.age += 20 too
	fmt.Println(lev.age)
}
