package main

import "fmt"

func main () {

	mySlice := []int {1,2,3,5,6,8,98,}  // length and capacity set to 7 integers

	for i, value := range mySlice {
		fmt.Println(i, "--- ",value)

	}
	fmt.Println(len(mySlice))
}
