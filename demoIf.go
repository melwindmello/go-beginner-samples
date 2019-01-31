package main

import "fmt"

func main(){

	//Demo if 
	var age int = 0
	fmt.Printf("Default age is %v\n", age)

	fmt.Println("Enter age:")
	fmt.Scanf("%d", &age)

	if age > 18  && age < 100 {
		fmt.Println("Eligible to vote")
	} else if age > 100 {
		fmt.Println("Invalid age")
	} else {
		fmt.Println("Not Eligible to vote")
	}

}