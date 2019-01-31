package main

import "fmt"


func main(){ 
	a, b := swap("melwin", "dmello")
	fmt.Println(a)
	fmt.Println(b)

	ans := addNumbers(2,3,4)
	fmt.Println(ans)
	ans = addNumbers(2,3,4,100)
	fmt.Println(ans)
	ans = addNumbers()
	fmt.Println(ans)

	fmt.Println("********* CLOSURE FUNCTION *************")
	closureFunctionDemo()

	fmt.Println("********* DEMO DEFER FUNCTION *************")
	deferFunctionDemo()

	//TODO : RECURVISE FUNCTION
}

func swap(fname string, lname string) (string, string ){
	fmt.Println("swap called.....")
	return lname, fname
}

func addNumbers(args ...int) int {
	fmt.Println("addNumbers called.....")
	total := 0

	for _,v := range args {
		//total = total + v
		total += v
	}

	return total
}

func closureFunctionDemo(){ 
	fmt.Println("closureFunctionDemo called.....")

	incrementBy1 := func (a int) int {
		return a + 1
	}

	fmt.Println(incrementBy1(1))
	fmt.Println(incrementBy1(11))
	fmt.Println(incrementBy1(100))

}

func deferFunctionDemo(){
	fmt.Println("deferFunctionDemo called.....")
	defer addNumbers(1,2)
	swap("melwin", "dmello")
	fmt.Println("deferFunctionDemo ends.....")
}