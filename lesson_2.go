package main

import (
	"fmt"
	"reflect"
)

func main() {
	//String 
	var message string
	message = "Data type String"
	fmt.Println(reflect.TypeOf(message))

	// Integers 
	var age int8
	var groups int16
	var card_balance int32
	var gold_card_balance int64

	age = 78
	groups = 20000
	card_balance = 2_000_000_000
	gold_card_balance = 9_223_372_036_854_775_807 // девять квинтильонов двести двадцать три квинтильона тридцать шесть квадриллионов восемьсот пятьдесят четыре триллиона семьсот семьдесят пять миллиардов восемьдесят семь.

	fmt.Println(age)
	fmt.Println(groups)
	fmt.Println(card_balance)
	fmt.Println(gold_card_balance)

	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(groups))
	fmt.Println(reflect.TypeOf(card_balance))
	fmt.Println(reflect.TypeOf(gold_card_balance))

	
}