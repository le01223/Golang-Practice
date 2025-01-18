package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}
type Airplace struct{}

func (c Car) move() {
	fmt.Println("Car is driving!!")
}
func (a Airplace) move() {
	fmt.Println("Airplane is flying!!")
}

func main() {
	boing_333 := Airplace{}
	honda := Car{}
	boing_333.move()
	honda.move()
}
