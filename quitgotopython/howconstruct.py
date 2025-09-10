def howconstruct(target,wordlist):
    if target=="":
        return []
    for x in wordlist:
        if target.startswith(x): 
            remainder=target[len(x):]
            newwordlist=howconstruct(remainder,wordlist)
            if newwordlist is not None:
                return newwordlist+[x]

    return None


    


print(howconstruct("oshan",["osh","an","a"]))
