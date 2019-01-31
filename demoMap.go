package main

import "fmt"

func main() {

	myMap := make(map[string]int)

	myMap["red"] = 100
	myMap["yellow"] = 200
	myMap["green"] = 300

	fmt.Println(myMap["red"])
	fmt.Println(myMap["yellow"])
	fmt.Println(myMap["green"])
	fmt.Println(myMap["green1"])
	fmt.Println(len(myMap))

}