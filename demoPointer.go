package main

import "fmt"

func main(){


	var num1 int
	num1 = 100
	num2 := num1
	
	var num3 *int
	num3 = &num1

	var nilPtr *int

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num3)

	num1 = 777
	num2 = 10
	fmt.Println("********************************")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num3)

	fmt.Println("********************************")
	fmt.Println(nilPtr)
	fmt.Println(*nilPtr) //panic. TODO handle this will recover
	

}