package main

import "fmt"

func main () {
	fmt.Println(greet("Mathew ", "Jonson "))
}

func greet (fname, lname string)(string, string) {
return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

// (string, string) defines the 'RETURNS' as two string types