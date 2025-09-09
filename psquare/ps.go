package main
func psq(num int){






}
func psqrec2(num int,array[]int)[]int{
	if num==0{
		return []int{}
	}
	if num<0{
		return nil
	}
	var smallpsq []int 
	
	for _,val:=range array{
		a:=psqrec(num-val,array)
		if a!=nil{
		 combo := append([]int{}, a...) // copy
            combo = append(combo, val)
		if len(combo)<len(smallpsq) || len(smallpsq)==0{
			smallpsq=a
			return smallpsq
		}
}


	}
return smallpsq
}
func psqrec(num int, array []int) []int {
    if num == 0 {
        return []int{}
    }
    if num < 0 {
        return nil
    }

    var smallpsq []int

    for _, val := range array {
        a := psqrec(num-val, array)
        if a != nil {
            combo := append([]int{}, a...) // copy
            combo = append(combo, val)

            if len(smallpsq) == 0 || len(combo) < len(smallpsq) {
                smallpsq = combo
            }
        }
    }
    return smallpsq
}

