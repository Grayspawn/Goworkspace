package main

import "fmt"

func main () {
	for l := 0; l <=100; l++ {
		if l%2 == 0 {
			fmt.Println(l)
		}
	}
}
