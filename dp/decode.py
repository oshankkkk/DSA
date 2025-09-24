def decode(stringnum,index):
    if index==5:
        return 1
    if stringnum[:index+1]=="0":
        return 0
    b=decode(stringnum,index+1)
    if int(stringnum[:index+2])<=26:
        b+=decode(stringnum,index+2)
    return  b
print(decode("11106",0))
