package main

import "fmt"

func main () {
 	fmt.Println(greet("Mick ", "Steve"))
}

func greet (fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
