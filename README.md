# DSA

## Dynamic programming
### UnderStanding bigO 
This space/time complexity is i think the resources yk memory and time your code uses at a time, ryt
#### Time complexity
so big O is ....
#### Space complexity
## Tips n Advice
when you see a new problem try n search for things you already know with
when a problem is given and if its big make it small by giving smaller values, cause its reduces the complexity
Okay Okay i think it kinda clicked lowkey for me, so dynamic programming problem is where a problem is when that problem is made up of several common subproblems.Greedy is differents it bout something bout doing the best solution on **each step**  and solving the problem.
In dynamic programming we talk bout subproblems not sub steps
So in dynamic programming there is no thinking in the code like unless its to end the recursion its we dont ryt if statements to solve the problem its just the same steps repeating.   

#### skipping the fibo problem cause thats too easy 

#### climbing stars problem and problem patterns in dynamic programming
def howconstruct(target,wordlist):
    if target=="":
        return []
    newwordlist=[]
    for x in wordlist:
        if target.startswith(x): 
            remainder=target[len(x):]
            newwordlist=howconstruct(remainder,wordlist)
            newwordlist.append(x)

    return newwordlist 

print(howconstruct("oshan",["osh","an","a"]))

