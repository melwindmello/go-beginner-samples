package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo go routine...")
	go forLoop1("T1")
	forLoop1("T2")
	//fmt.Println("Demo go routine...")
}

func forLoop1(varName string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("varName = %s, i = %d\n", varName, i)
		time.Sleep(time.Second * 1)
	}
}
