package main

import (
	"fmt"
	"net/http"
)

/*	func main() {
	fmt.Println("It is currently:", time.Now())
	fmt.Println("A random number from 1 - 100", rand.Intn(100))
}*/

/*	func add (x, y int) int {
		return x + y
	}

	func main () {
		fmt.Println(add(53, 79))
	}*/

/*	#multiple-results.go
	func swap (x, y string) (string, string){
		return y, x
	}

	func main () {
		a, b := swap("fucker", "mother")
		fmt.Println(a, b)

	}*/

/*	#named-results.go
	func split(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	func main() {
		fmt.Println(split(17))
	}*/

/*	#variables.go
	var c, python, java bool

	func main () {
		var i int
		fmt.Println(i, c, python, java)

	}*/

/*	#variables-with-initialisers.go
	 var i, j int = 1, 2

	func main () {
		var c, python, java = true, false, "no!"
		fmt.Println(i, j, c, python, java)
	}*/

/*	#short-variable-declarations.go
	func main () {
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no fucking way"
		fmt.Println(i, j, k, c, python, java)
		fmt.Println(python, java, j, c, i)
	}*/

/*	#basic-types.go
	var (
		ToBe bool = false
		MaxInt uint64 = 1<<64 -1
		z complex128  = cmplx.Sqrt(-5 + 12i)
	)

	func main() {
		fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", z, z)
	}*/

/*	#zero.go
	func main () {
		var i int
		var f float64
		var b bool
		var s string

		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	}*/

/*	#type-conversions.go
	func main () {
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x+ y*y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	}*/

/*#book chapetr 3 types
func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hellow World"[2])
	fmt.Println("Hello" + "World")
}*/

/*	#book chapetr 3 types
	func main () {
		fmt.Println(true && true)
		fmt.Println(true && false)
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(!false)
	}*/

/*	#constants.go
	const Pi = 3.14

	func main () {
		const World = "FUCK YEAH"
 		fmt.Println("America.....", World)
		fmt.Println("Happy", Pi, "Day")

		const Truth = true
		fmt.Println("Go has rules?", Truth)
	}*/

/*#for.go
func main () {
	sum :=0
	for i := 0; i < 10; i++ {
		sum += i
	}
fmt.Println(sum)
}*/

/*	#operators
	func main () {
		 p := 10
		 q := 2
		p /= q
		fmt.Println(p)
	}*/

/* #forloop.go
func main () {
		sum := 1
		for sum <1000 {
			sum += sum
		}

		fmt.Println(sum)
	}*/

/*#programmingpythn.net
func main () {
	fmt.Println("A number frm 1 -100", rand.Intn(50))
}*/

/*func add (x, y float64) float64 {
	return x+y
}*/

/*#go language types (programmingpython.net)
func main () {
	num1, num2 := 5.6, 9.5

	fmt.Println(add(num1,num2))

}*/

/*func multiple (a, b string) (string, string) {
	return a,b
}

func main () {
	w1, w2 := multiple("Hey", "there")
	fmt.Println(w1, w2)
}*/

/*#programmingpython.net
func main () {
	var a int = 54
	var b float64 = float64(a)

	x := a // x will be type int
}*/

/* pointers (programmingpython.net)
func main () {
	x := 15
	a := &x // points to x (memory address)

	fmt.Println(a) // prints our memory address
	fmt.Println(*a) // read a through the pointer so this will print value of 15

	*a = 5 // sets the value pointed at to 5, so now x equals 5
	fmt.Println(x) // see new value of x -5

	*a= *a**a // assign new value to a. What is the value of x?
	fmt.Println(*a) // prints a value
	fmt.Println(x) // prints a value

	fmt.Println(&x) // prints a memory address
	fmt.Println(a) // prints a memory address

} */

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is supposed to be neat!")
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert Go and Blockchain Programmer Leslie Onyesoh!")
}

func main() {
	http.HandleFunc("/", indexhandler) // path, then what function to run
	http.HandleFunc("/about/", abouthandler)
	http.ListenAndServe(":8000", nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now
}
