package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gettinguserinput()int{
fmt.Println("give the amount of stairs")
reader:=bufio.NewReader(os.Stdin)
userinput,err:=reader.ReadString( '\n')
if err!=nil{
fmt.Println(err)

}
userinput=strings.TrimSpace(userinput)
usernumber,err:=strconv.Atoi(userinput)
if err!=nil{
fmt.Println(err)

}
return usernumber
}
func main  (){

	fmt.Println("question1")
	a:=costclimbing(2,[]int{1,100,1,1,1,100,1,1,100,1})
	fmt.Println(a)

}
