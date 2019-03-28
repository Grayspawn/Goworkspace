<<<<<<< HEAD
 package main

 import "fmt"

 const (

 	_ = iota
 	KB = 1 << (iota * 10) // 1 << (1 * 10)
 	MB = 1 << (iota * 10) // 1 << (2 * 10)
 	GB = 1 << (iota * 10) // 1 << (3 * 10)
 	TB = 1 << (iota * 10) // 1 << (4 * 10)
 )

func main () {
	fmt.Println("binary \t \t decimal")
	fmt.Printf("%b \t", KB)
	fmt.Printf("%d \n", KB)
	fmt.Printf("%b \t", MB)
	fmt.Printf("%d \n", MB)
	fmt.Printf("%b \t", GB)
	fmt.Printf("%d \n", GB)
	fmt.Printf("%b \t", TB)
	fmt.Printf("%d \n", TB)



}
=======
package _7_iota
>>>>>>> da815ba802ee2d58bcda9cb78566aa0020bf85a9
