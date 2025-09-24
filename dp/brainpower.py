def brainpower(qlist,index):
    if index>=len(qlist):
        return 0
    ans=qlist[index][0]+brainpower(qlist,index+qlist[index][1]+1)
    ans2=brainpower(qlist,index+1)

    return max(ans,ans2)

questions = [[3,2],[4,3],[4,4],[2,5]]
print(brainpower(questions,0))  #5

