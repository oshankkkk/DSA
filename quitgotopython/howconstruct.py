def howconstruct(target,wordlist):
    if target=="":
        return []
    bhiglist=[]
    for x in wordlist:
        if target.startswith(x): 
            remainder=target[len(x):]
            newwordlist=howconstruct(remainder,wordlist)
            return newwordlist+[x]

    return bhiglist


    

print(howconstruct("oshan",["osh","an","a"]))
