package main


func climbingstars(amountstairs int)int{
 if amountstairs==0{
return 1

 } 

 if amountstairs<0{
return 0
 }

 return climbingstars(amountstairs-1)+ climbingstars(amountstairs-2)
}
