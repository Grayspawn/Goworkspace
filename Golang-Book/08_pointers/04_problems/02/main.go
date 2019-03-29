package main

import "fmt"

func swap (xPtr, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr
}

func main () {
	x, y := 1, 2
	fmt.Println(x,y)
	swap(&x, &y)
	fmt.Println(x,y)

}