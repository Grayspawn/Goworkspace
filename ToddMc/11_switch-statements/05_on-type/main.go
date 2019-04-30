package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchonType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchonType(7)
	SwitchonType("Onyesoh")
	var t = contact{"Good to see you,", "Leslie"}
	SwitchonType(t)
	SwitchonType(t.greeting)
	SwitchonType(t.name)

}
