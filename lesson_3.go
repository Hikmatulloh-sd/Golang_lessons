package main

import "fmt"

// Область видемости 

var a,b,c int 

func main(){
	a,b,c = 1,2,3

	fmt.Println(a,b,c)
	print()
}

func print(){
	// a,b,c := 4,5,6

	fmt.Println(a,b,c)
}