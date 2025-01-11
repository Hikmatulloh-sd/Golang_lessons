package main

import (
	"fmt"
)

func main(){
	// Scan	
	var name string
	var age int 

	fmt.Println("Input your name: ")
	fmt.Scan(&name)

	fmt.Println("Input your age: ")
	fmt.Scan(&age)

	fmt.Printf("Name: %s, Age: %d \n",name,age)

	var a,b int

	fmt.Println("Enter numbers: ")
	fmt.Scanln(&a,&b)

	sum := a + b
	fmt.Println(sum)
	
	// // Арифметические операции 
	// var a,b int = 10, 6

	// sum := a + b
	// differ := a - b
	// multi := a * b
	// // num := a / b // Деление без остатка 
	// num := a % b // Остаток от делении
	
	// a++ // a = a + 1 (--)

	// fmt.Println(sum)
	// fmt.Println(differ)
	// fmt.Println(multi)
	// fmt.Println(num)
	// fmt.Println(a)
}