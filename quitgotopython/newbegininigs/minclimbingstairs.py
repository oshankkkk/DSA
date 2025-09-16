def climbingstairsmin(costlist,lenofcost):
    if lenofcost==0:
        return costlist[lenofcost]
    if lenofcost==1:
        return costlist[lenofcost]

    return costlist[lenofcost]+ min(climbingstairsmin(costlist,lenofcost-1),climbingstairsmin(costlist,lenofcost-2))


costlist=[10,15,20]
lenofcost=len(costlist)
print(min(climbingstairsmin(costlist,lenofcost-1),climbingstairsmin(costlist,lenofcost-2)))

