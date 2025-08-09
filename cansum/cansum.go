package main

import (
	"fmt"
)

func main (){

	memo:=make(map[int]bool)
//	array:=[]int{7,14}
fmt.Println(countsum(300, []int{7,14}, memo)) // prints: true
fmt.Println(countsum(301, []int{7,14}, memo)) // prints: false (because 301 can't be made from 7 and 14)

}
func countsum(num int, array []int,memo map[int]bool)bool {
	_,exits:=memo[num]
	if exits{
return memo[num]
	}
if num==0{

	return true
}

if num<0{
return false
}
for _,value:=range array{
	a:= num-value
	if countsum(a,array,memo){

memo[num]=true
return memo[num]
	}
}
memo[num]=false
// final return after all recursion
return memo[num]

//return false

}








