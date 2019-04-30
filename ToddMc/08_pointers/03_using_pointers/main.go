package main

import "fmt"

func main() {

	a := 3993

	fmt.Println(a)  // 3993
	fmt.Println(&a) // 0xc000080008

	var b = &a
	fmt.Println(b)  // 0xc000080008
	fmt.Println(*b) // 3993

	*b = 4999      // changes value at that memory address and so changes value of a
	fmt.Println(a) // 4999
}

/*we can pass a memory address instead of a bunch of values (we can pass a reference)
and then we can still change the value of whatever is stored at that memory address
this makes our programs more performant.

we don't have to pass around large amounts of data, we only have to pass around addresses

Everything is PASS BY VALUE in go, btw
when we pass a memory address, we are passing a value*/
