package main

import "fmt"

func main() {
	switch "Steve" {
	case "Stanley":
		fmt.Println("Hi Stanley")
	case "Michael":
		fmt.Println("Hi Michael")
	case "Steve":
		fmt.Println("Steve ....you de man!!")
        fallthrough //  On match to true the fallthrough statement automatically /prints the next case clause as true. The expression on the next case is not considered when transferring control
	case "Pete":
		fmt.Println("Hi Pete")
		fallthrough
	case "Aereon":
		fmt.Println("Hi Aereon")
	}
}