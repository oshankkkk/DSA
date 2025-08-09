package main

import "fmt"
func main(){
	grid:=make(map[string]int)
	a:=gridtraveler(3,3,grid)

	b:=gridtraveler(18,18,grid)
fmt.Println(a)

fmt.Println(b)
}
func gridtraveler(row,col int,grid map[string]int)int{
	key:=fmt.Sprintf("%d , %d",row,col)
	num,exits:=grid[key]
	if exits{
return num
	}
if row==0 || col ==0{
return 0
}
if row==1 && col ==1{
return 1}else{
grid[key]=gridtraveler(row-1,col,grid)+gridtraveler(row,col-1,grid)


}
return grid[key]
}
