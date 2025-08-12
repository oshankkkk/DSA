package main

func climbingstairsmemo(amountstairs int,memo map[int]int)int{
	if value,exits:= memo[amountstairs]; exits{
return value
	}
 if amountstairs==0{
return 1

 } 

 if amountstairs<0{
return 0
 }

 memo[amountstairs]=climbingstars(amountstairs-1)+ climbingstars(amountstairs-2)
 return memo[amountstairs]
}













