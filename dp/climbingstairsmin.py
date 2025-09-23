def climbingmin(staircount,costlist):
    if staircount<0:
        return 0
    if staircount==0 or staircount==1:
        return costlist[staircount]
    return costlist[staircount-1] + min(
        climbingmin(staircount-1,costlist),
        climbingmin(staircount-2,costlist)
    ) 

print(min(
    climbingmin(6-1,[10,1,10,1,10,1]),
    climbingmin(6-2,[10,1,10,1,10,1])
))

