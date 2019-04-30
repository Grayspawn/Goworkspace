package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println(x)
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

}
