package main
func cbs(staircount int)int{
if staircount==0{
	return 1
}
if staircount<0{
	return 0
}
return cbs(staircount-1)+cbs(staircount-2)
}
