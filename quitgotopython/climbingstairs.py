def climbstairspy(stairlist,target):
    if target==0:
        return stairlist[target]
    if target<0:
        return None
    min=0
    for i in range (1,3):
        a=climbstairspy(stairlist,target-i)
        if a is not None:
            a+=stairlist[target]
            if min==0 or min>a:
                min=a
    return min

 





def minCostClimbingStairsRecursive(cost):
    n = len(cost)

    def dfs(i):
        if i >= n:
            return 0
        return cost[i] + min(dfs(i+1), dfs(i+2))

    return min(dfs(0), dfs(1))
g=minCostClimbingStairsRecursive([10,15,20])
print(g)
