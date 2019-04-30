package main

import "fmt"

func main (){
   myFriendsname := "Ted" // changing name length to 2 will print the first line

	switch  {
	case len(myFriendsname) == 2:
		fmt.Println("Wassup my fiend with the name of length 2")
	case myFriendsname == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsname == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsname == "Teddie", myFriendsname == "Teddy":
		fmt.Println("Your name is either Teddie or Teddy")
	case myFriendsname == "Julian":
		fmt.Println("Wassup Julian")
	default:
		fmt.Println("Nothing matched, this is the default")



	}

}
