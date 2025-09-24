def uniquepath(rows,cols):
    if rows==1 and cols ==1:
        return 1
    if rows==0 or cols==0:
        return 0

    return uniquepath(rows-1,cols)+uniquepath(rows,cols-1)
     

