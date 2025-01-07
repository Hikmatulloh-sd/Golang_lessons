package main

import (
	"fmt"
)

// variables | переменные 
// Base data types | Основные типы данных 

func main(){
	// Integers Целое число
	var moth int8 = 8 
	var tickets int16 = 20000
	var	blue_card int32 = 2_000_000_000
	var gold_card int64 = 1_000_000_000_000

	fmt.Println(moth)
	fmt.Println(tickets)
	fmt.Println(blue_card)
	fmt.Println((gold_card))

	// Uint >= 0
	var number uint8 = 34
	fmt.Println(number)

	// String Сторока
	hello := "Hello world"
	fmt.Println(hello)
	fmt.Println(string(97))

	// Float Веществонное число 
	var a float32 = 1.5
	var b float64 = 1.55555

	fmt.Println(a,b)
}	


// Целочисленные типы данных
// Тип	Описание	Диапазон чисел
// int8	8-битные числа со знаком.	от -128 до 127
// int16	16-битные числа со знаком.	от -32,768 до 32,767
// int32	32-битные числа со знаком.	от -2^31 до 2^31-1
// int64	64-битные числа со знаком.	от -2^63 до 2^63-1