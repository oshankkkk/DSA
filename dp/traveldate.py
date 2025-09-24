def ticketday(days,costlist,index):
    if index==len(days):
        return 0
    a=costlist[0]+ticketday(days,costlist,index+1)
    j=index
    while j<len(days) and days[j]<days[index]+7:
        j+=1

    b=costlist[2]+ticketday(days,costlist,index+7)

    j=index
    while j<len(days) and days[j]<days[index]+30:
        j+=1

    c=costlist[2]+ticketday(days,costlist,index+30)

    return min(a,b,c)

    
