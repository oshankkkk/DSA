package main
func psq(num int){




	


}
func psqrec(num int)[]int{
	array:=make([]int,1)
	if num==0{
		return []int{}
	}
	if num<0{
		return nil
	}

	var smallpsq []int 
	for i:=1;i<num;i++{
		if i*i>=num{
			break
		}else{
			array=append(array,i)
		}
	}
	for _,val:=range array{
		a:=psqrec(num-val)
		if a!=nil{
		a=append(a, val)
		if len(a)<len(smallpsq) || len(smallpsq)==0{
			smallpsq=a
		}
}


	}
return smallpsq
}
