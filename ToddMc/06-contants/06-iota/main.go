package main

import "fmt"

const (

	_ = iota // 0  - throw away the 0 - we are not going to use it
	b = iota * 10  // 1*10
	c = iota * 10  // 2*10

)

func main () {
	fmt.Println(b)
	fmt.Println(c)

}
