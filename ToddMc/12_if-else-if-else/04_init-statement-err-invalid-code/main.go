package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	fmt.Println(food) //Will not run as this is outside the scope of the parameter food
}
