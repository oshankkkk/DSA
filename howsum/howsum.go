package main 
import "fmt"

func main(){
	fmt.Println("how some boys")
	memo:=make(map[int][]int)
	list:=howsum(7,[]int{2,4},memo)
	fmt.Println(list)
}
func howsum(num int,array []int,memo map[int][]int)([]int){
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
for _,values:=range array{
	remainder:=num-values
	list:=howsum(remainder,array,memo)
	if list!=nil{
		memo[num]=append(list,values )
return memo[num]
	}

}
memo[num]=nil

return memo[num]
}

