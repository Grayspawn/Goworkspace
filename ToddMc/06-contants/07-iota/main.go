<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> cadc18a0c75a281ba9cd3ee75eaf29ba563c7f8d
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



<<<<<<< HEAD
}
=======
}
=======
package _7_iota
>>>>>>> da815ba802ee2d58bcda9cb78566aa0020bf85a9
>>>>>>> cadc18a0c75a281ba9cd3ee75eaf29ba563c7f8d
