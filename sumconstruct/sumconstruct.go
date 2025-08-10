package main

import (
	"fmt"
	"strings"
)
func main(){
	fmt.Println("hi:")
	memo:=make(map[string]int)
	a:=sumconstruct("abcd",[]string{"gu","yy","kj"},memo)
	fmt.Println(a)
}

func sumconstruct(word string,array[]string,memo map[string]int)int{
if word==""{
	memo[word]=1
return memo[word]
}
var count int
for _,cands:=range array{
if strings.HasPrefix(word,cands){
	remainder:=word[len(cands):]
	count+=sumconstruct(remainder,array,memo)
	memo[word]=count


}

}
return memo[word]
}
