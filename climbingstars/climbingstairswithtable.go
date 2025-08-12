package main




func climbingstairstable(num int)int{
	staircase:=make([]int,num+1)
staircase[0]=1

staircase[1]=1
	for i:=2;i<=num;i++{
		staircase[i]+=staircase[i-1]+staircase[i-2]
		

	}
	
return staircase[num]



}
