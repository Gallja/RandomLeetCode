def removeDuplicates(nums: List[int]) -> int:
    cnt_duplicates = 0
    k = 0
    last_num = 200
    
    for i in range(len(nums)):
        for j in range(i+1, len(nums)):
            if nums[i] == nums[j] and last_num != nums[i]:
                last_num = nums[i]
                cnt_duplicates += 1
                break
            elif nums[i] == nums[j] and last_num == nums[i]:
                cnt_duplicates += 1

    k = len(nums) - cnt_duplicates
    return k

nums_1 = [1, 1, 2]
nums_2 = [0,0,1,1,1,2,2,3,3,4]

print(removeDuplicates(nums_1), "\n")
print(removeDuplicates(nums_2))