package main

import "fmt"

func main () {
	m := make([]string, 1,25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m) // Ricky
}

func changeMe (z []string){
	 z[0] = "Ricky"
	 fmt.Println(z) // Ricky
}
