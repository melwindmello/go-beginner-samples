package main

import "fmt"

func main() {

	var n1, n2 int
	n1 = 10
	n2 = 0

	if n2 == 0 {

		defer func () {
			str := recover()
			fmt.Println(str)
		}()

		panic("Div by 0") 

	} else {
		fmt.Println(n1/n2)
	}

}
