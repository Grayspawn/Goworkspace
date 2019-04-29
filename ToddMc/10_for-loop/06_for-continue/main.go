package main

import "fmt"

func main () {
	i :=0
	for {
		i++ // ++ incrementer operand adding 1. It is not a 'post' as it executes earlier
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
