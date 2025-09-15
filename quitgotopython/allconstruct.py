def allconstructpy(target,wordlist):
    if target=="":
        return [[]]
    a=[]
    for x in wordlist:
        if target.startswith(x):
            remainder=target[len(x):]
            b=allconstructpy(remainder,wordlist)
            for i in b:
                a.append([x]+i)
    return a
            




print(allconstructpy("oshan",["oshan","o","sh","a","n","osh","an","2","Ss","EDD"]))



