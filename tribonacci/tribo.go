package main

func tribo(num int)int{
	array:=make([]int,num+1)
	array[1]=1
	array[2]=1

	for i:=3;i<=num;i++{
		array[i]=array[i-1]+array[i-2]+array[i-3]
	

	}
return array[num]
}
