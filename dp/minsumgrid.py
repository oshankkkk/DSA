def minsumgrid(grid,currentrow=0,currentcol=0):
    if currentrow>=len(grid) or currentcol>=len(grid[0]):
        return float("inf") 
    if currentrow==len(grid)-1 and currentcol>=len(grid[0])-1:
        return grid[currentrow][currentcol]
    return grid[currentrow][currentcol]+min(minsumgrid(grid,currentrow+1,currentcol),minsumgrid(grid,currentrow,currentcol+1))
    
grid = [
    [1,3,1],
    [1,5,1],
    [4,2,1]
]
print(minsumgrid(grid))  

