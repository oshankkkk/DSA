package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func febonacci(num int)int{

if num==0 || num==1{
return num

}else{

return febonacci(num-2)+febonacci(num-1)
}


}
func main(){

fmt.Println("febonacci in recursion")
reader:=bufio.NewReader(os.Stdin)
input,_:=reader.ReadString('\n')
input=strings.TrimSpace(input)
num,_:=strconv.Atoi(input)

	memo:=make(map[int]int)
fmt.Println(febowithmemo(num,memo))
} 

func febowithmemo(num int, memo map[int]int)int{
	value, exits:=memo[num]
	if exits{
return value
	}else 
if num==0 || num==1{
return num

}else{
memo[num]=febowithmemo(num-2,memo)+febowithmemo(num-1,memo)

return memo[num]
}

 }
