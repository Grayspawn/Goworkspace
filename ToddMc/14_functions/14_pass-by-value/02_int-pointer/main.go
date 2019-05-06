package main

import "fmt"

func main () {

 age := 54
 fmt.Println(age)
 fmt.Println(&age) // print memory address/location


 changeMe(&age)

 fmt.Println(age)
 fmt.Println(&age)
}

func changeMe (z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 34  // Assign 34 to the value of the address
	fmt.Println(z)
	fmt.Println(*z)
}
