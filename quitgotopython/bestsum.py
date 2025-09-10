def bestsum(target,numlist):
    if target==0:
       return []  
    if target<0:
        return None
    shorttestlist=[] 
    for x in numlist:
        sumlist=bestsum(target-x,numlist)
        if sumlist!=None:
            sumlist.append(x)
            if len(sumlist)<len(shorttestlist) or len(shorttestlist)==0:
                shorttestlist=sumlist
    return shorttestlist



a=bestsum(7,[3,4,5,1,7])
print(a)

