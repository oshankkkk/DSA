def allconstructclone1(target,wordlist):
    if target =="":
        return [[]] 
    mylist=[]
    for word in wordlist:
        if target.startswith(word):
            remainder=target[len(word):]
            mylist2=allconstructclone1(remainder,wordlist)
            for elemant in mylist2:
                elemant.append(word)
                mylist.append(elemant)

    return mylist

def bestsumclone(target,numlist):
    if target==0:
        return[]
    if target<0:
        return None
    minlist=[]
    for num in numlist:
        remiander=target-num
        a=bestsumclone(remiander,numlist)
        if a is not None:
            a.append(num)
            if len(minlist)>len(a) or len(minlist)==0:
                minlist=a

    return minlist

#print(allconstructclone1("oshan",["oshan","o","sh","a","n","osh","an","2","Ss","EDD"]))

a=bestsumclone(7,[3,4,5,1,7])
#print(a)


def allconstructtableclone(target,wordlist):
    targetcount=len(target)+1
    newlist=[[] for _ in range(targetcount)]

    newlist[0]=[[]]
    for x in range(targetcount) :
        remainder=target[x:]
        for word in wordlist:
            if remainder.startswith(word):
                for i in newlist[x]:
                    newlist[x+len(word)].append(i+[word])


    return newlist[len(target)]






print(allconstructtableclone("oshan",["oshan","o","sh","a","n","osh","an","2","Ss","EDD"]))






