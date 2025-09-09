package main

import "fmt"

//currentstair:=0
func costclimbing(currentstair int,stairs []int)int{
//	var numstairs= len(stairs) 
	if currentstair<0{
return 0
	}
var a int	
	for i:=1;i<=2;i++{


		a = costclimbing(currentstair-i,stairs)

		a+=stairs[currentstair]
		fmt.Println(a)

	}
return a

}
