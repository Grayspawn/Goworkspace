package main

import "fmt"

type customer struct {
	name string
	age int
}

func main () {
	c1 := customer{"Leslie", 54}
	fmt.Println(&c1.name)

	changeMe(&c1)

	fmt.Println(c1)
	fmt.Println(&c1.name)
}

func changeMe (z *customer) {
	fmt.Println(z)
	fmt.Println(&z.name)
	z.name = "Spicy"
	fmt.Println(z)
	fmt.Println(&z.name)
}
