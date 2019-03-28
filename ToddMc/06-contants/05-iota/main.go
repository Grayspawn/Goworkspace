package main

import "fmt"

const (

<<<<<<< HEAD
	a = iota // 0
	b = iota // 1
	c = iota // 2
=======
	a =iota
	b = iota
	c = iota
>>>>>>> da815ba802ee2d58bcda9cb78566aa0020bf85a9

)

const (

<<<<<<< HEAD
	d = iota // 0
	e = iota // 1
	f = iota // 2
=======
	d =iota
	e = iota
	f = iota
>>>>>>> da815ba802ee2d58bcda9cb78566aa0020bf85a9

)

func main () {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

