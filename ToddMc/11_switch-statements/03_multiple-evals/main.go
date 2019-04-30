package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Benny":
		fmt.Println("Wassup Tim, & Benny ")
	case "Mark", "Jenny":
		fmt.Println("Hey there Mark & Jenny")
	case "Pete", "jenny":
		fmt.Println("Wassup Pete and little jenny")
	}
}
