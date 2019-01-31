package main

import "fmt"

func main(){

	var i int
	for i=0; i<100; i++ {
		fmt.Println(i)
		if i == 12 {
			break
		}
	}

	fmt.Println("END.")

}