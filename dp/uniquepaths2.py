def uniquepath2(grid,currentrow,currentcol):
    lenrow=len(grid)
    lencol=len(grid)

    if grid[currentrow][currentrow]==1 or currentrow>=lenrow or currentcol>=lencol:
        return 0
    if currentrow==lenrow-1 & currentcol==lencol-1:
        return 1
    return uniquepath2(grid,currentrow-1,currentcol)+uniquepath2(grid,currentrow,currentcol+1)


    


        
         









