//TODO

package main

import "fmt"

func main()  {

  /*var x = make([]int,5)
  fmt.Printf("%v",x)
  var arr = []int{1,2,3,4,5}
  s:=arr[-1:3]
  fmt.Printf("%v",s)*/
  var x=[]int{1,2,3,4,5,6,7}
  sarr:= x[1:4]
  fmt.Printf("size:%d, capacity:%d\n\n",len(sarr),cap(sarr))
  slice1:=append(x,8,9)
  fmt.Printf("%v\n",slice1)
  slice2:=make([]int,2)
  copy(slice2,x)
  fmt.Println(x,slice2)
}
