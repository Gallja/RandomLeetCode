def canPartitionGrid(grid) -> bool:
    total_sum = sum(sum(row) for row in grid)
    if total_sum % 2 != 0:
        return False
    
    target = total_sum // 2

    curr_sum = 0

    for i in range(len(grid)):
        curr_sum += sum(grid[i])
        if curr_sum == target:
            return True
        
    curr_sum = 0
    
    for i in range(len(grid[0])):
        curr_sum += sum(grid[j][i] for j in range(len(grid)))
        if curr_sum == target:
            return True

    return False

grid_1 = [[6, 1, 4], [2, 5, 4]]
print(canPartitionGrid(grid_1)) #true

grid_2 = [[1,4],[2,3]]
print(canPartitionGrid(grid_2)) #true

grid_3 = [[1,3],[2,4]]
print(canPartitionGrid(grid_3)) #false