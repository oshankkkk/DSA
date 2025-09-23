def robhouse(housecount,houselist):
    if housecount ==1 or housecount==0:
        return houselist[housecount]
    if housecount<0:
        return 0
    return houselist[housecount]+ robhouse(housecount-2,houselist)
houselist= [2,7,9,3,1]
housecount=len(houselist)
print(max(robhouse(housecount-1,houselist),robhouse(housecount-2,houselist)))
