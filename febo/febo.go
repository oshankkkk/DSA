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

fmt.Println(febonacci(num))
} 


