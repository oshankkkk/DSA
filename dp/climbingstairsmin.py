def climbingmin(staircount,costlist):
    if staircount==0 or staircount==1:
        return costlist[staircount]
    return costlist[staircount-1]+ min(climbingmin(staircount-1,costlist),climbingmin(staircount-2,costlist)) 

print(min(climbingmin(3-1,[10,15,20]),climbingmin(3-2,[10,15,20])))



    
    






