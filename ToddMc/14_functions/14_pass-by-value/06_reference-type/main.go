package main

import "fmt"

func main () {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Rocky"])
}

func changeMe(z map [string]int) {
	z["Rocky"] = 54
}
