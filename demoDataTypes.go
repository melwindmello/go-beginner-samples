package main

import "fmt"

func main(){

	fmt.Println("Datatypes...")

	var num1 = 10;
	fmt.Println("num1 = ", num1)

	fmt.Println("Value of num1=%d", num1)
	fmt.Printf("Value of num1=%d", num1) //Printf used for formatting

	//Type inference operator bein used :=
	myvar := true
	fmt.Printf("The datatype of myvar is %T\n", myvar) //%T show datatype

	//mutiple varibales declaration
	var(
		a = 10
		b = 20.5
		c = false
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//constants have to be declared and assigned a value
	const pi float32 = 3.14
	fmt.Println(pi)

	//mutiple variable declartion of different types
	var n1, n2, n3, n4 = 2, 3.5, "hello", false
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	

} //end of main()