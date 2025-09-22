def perfectsquare(number):
    if number==0 or number==1:
        return number
    count=float("inf")
    for i in range(number,0,-1):
        square=i*i
        if square<=number:
            a=perfectsquare(number-square)
            count = min(count, 1 + a)
    return count

b=perfectsquare(9)
print(b)









