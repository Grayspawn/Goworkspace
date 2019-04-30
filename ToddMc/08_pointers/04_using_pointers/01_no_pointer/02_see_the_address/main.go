package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in func main
	fmt.Println(&x)        // address in func zero
	zero(x)
	fmt.Println(x) // x is still equal to 5

}
