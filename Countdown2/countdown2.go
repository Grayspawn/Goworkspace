package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Happy New Year!!")

	// }
	// fmt.Println("It is currently:", time.Now())
	// fmt.Println("A number from 1 -100", rand.Intn(100))
}
