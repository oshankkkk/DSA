package main

import "fmt"
func main(){
//	grid:=make(map[string]int)
//	a:=gridtraveler(3,3,grid)

//	b:=gridtraveler(18,18,grid)
//fmt.Println(a)

//fmt.Println(b)
gridtravelertabulation(3,3)
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
func gridtravelertabulation(rows,cols int){
	grid:=make([][]int,rows+1)
	for index,_:=range grid{
		grid[index]=make([]int,cols+1)
	}
	for index,_:=range grid{
		for innerindex,_:=range grid[index]{
			grid[index][innerindex]=0
		}

	}


grid[1][1]=1
for i:= 1;i<=rows;i++{
	for j:=1;j<=cols;j++{
		if i==1 && j==1{
continue
		}
	grid[i][j]+=grid[i-1][j]+grid[i][j-1]



	}

}





	for _,value:=range grid{
		for _,innerval:=range value{
			fmt.Print(innerval)
		}
		fmt.Println("")

	}

}
