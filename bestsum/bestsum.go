package main

import "fmt"

func main() {
	fmt.Println("best some boys")
	memo := make(map[int][]int)
 sp := bestsum(100, []int{1,2,5,25},memo)
	fmt.Println( sp)

}

func bestsum(num int,array []int,memo map[int][]int)([]int){
	value,exists:= memo[num] 
	if exists{
		return value
	}

if num==0{
return []int{}
}
if num<0{
return nil
}

var shortnestpath []int

for _,values:=range array{

	remainder:=num-values
	list:=bestsum(remainder,array,memo)
	if list!=nil{
		list=append(list,values )
	if len(list)<len(shortnestpath)||  len(shortnestpath)==0{

shortnestpath=list
memo[num]=shortnestpath
	} }

	}



return memo[num]
}

