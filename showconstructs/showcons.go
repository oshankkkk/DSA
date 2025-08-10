package main

import (
	"fmt"
	"strings"
)
func main(){
fmt.Println("show combinations")
memo := make(map[string][][]string)
a:=showcons("abcd",[]string{"adbc","ab","a","abc","d","c","b"},memo)
fmt.Println(a)
for _,t:=range a{
fmt.Println(t)
}
}
func showcons(word string,cands[]string, memo map[string][][]string )[][]string{
if word==""{
memo[word]=[][]string{{}}
return memo[word]
}

var biglist [][]string
for _,cand:=range cands{
if strings.HasPrefix(word,cand){
	remainder:=word[len(cand):]
	list :=showcons(remainder,cands,memo)
	for _, combo := range list{
				newCombo := append([]string{cand}, combo...)
				biglist= append(biglist, newCombo)
				memo[word]=biglist
			}
}
} 
return memo[word]
}
