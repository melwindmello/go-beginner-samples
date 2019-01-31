package main

import "fmt"

func main(){

	var employeeList[4] string 

	employeeList[0] = "melwin"
	employeeList[1] = "melwin1"
	employeeList[2] = "melwin2"
	employeeList[3] = "melwin3"


	for i:=0; i<4; i++ {
		fmt.Printf("employeeList[%d] = %s\n", i, employeeList[i])
	}

	var twoDimArray = [2][2] int {{1,2}, {3,4}}

	for i:=0; i<2; i++ {
		for j:=0; j<2; j++ {
			fmt.Printf("twoDimArray[%d][%d] = %d\n", i, j, twoDimArray[i][j])
		}
	}
	

}