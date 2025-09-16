def climbingstairs(count):
    if count==0:
        return 1
    if count <0:
        return 0
    return climbingstairs(count-1)+climbingstairs(count-2)
print(climbingstairs(3))
    

