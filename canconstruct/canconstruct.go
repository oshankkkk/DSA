package main

import (
	"fmt"
	"strings"
)
func main(){
fmt.Println("can word make")
memo:=make(map[string]bool)
a:=canword("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeffff",[]string{"eeeeeeeeee","eeeeeeeeeeeeeeeeeeeeeeee","eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee","eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},memo)
fmt.Println(a)
}

func canword(word string, wordbank []string,memo map[string]bool)bool{
	if value,exits:=memo[word]; exits{
	return value
	}


if word==""{
memo[word]=true
return memo[word]
}
// for _,value :=range wordbank{
//	 if strings.Contains(word,value){
//	remainder:=strings.Split(word,value)
//fmt.Println(remainder[0],"here look")
//	memo[word] = canword(remainder[0],wordbank,memo)
//	return memo[word]
//	 }
//
//
//
//
//}
for _, w := range wordbank {
        if strings.HasPrefix(word, w) {
            remainder := word[len(w):]
            if canword(remainder, wordbank, memo) {
                memo[word] = true
                return true
            }
        }
    }
memo[word]=false
return memo[word]
}







