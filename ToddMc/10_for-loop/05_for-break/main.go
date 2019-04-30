package main

import "fmt"

func main() {
	i := 0 // initialisation
	for {
		fmt.Println(i)
		if i >= 10 { // condition
			break
		}
		i++ // post
	}
}
