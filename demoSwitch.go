package main

import "fmt"

func main(){

	//Demo if 
	var dow int
	fmt.Println("Enter Day of Week:")
	fmt.Scanf("%d", &dow)

	switch dow {
		case 1:
			fmt.Println("Sunday")	
		case 2:
			fmt.Println("Monday")
		case 3:
			fmt.Println("Tuesday")
		case 4:
			fmt.Println("Wednesday")
		case 5:
			fmt.Println("Thursday")
		case 6:
			fmt.Println("Friday")
		case 7:
			fmt.Println("Saturday")	
		default:
			fmt.Printf("Invalid entry: %v", dow)	

	}

	
}