def removeDuplicates(nums: List[int]) -> int:
    if not nums:
        return 0
    
    cnt_duplicates = 0
    k = 0
    last_num = 200

    array_1_indexes = []
    array_2_indexes = []

    flag = True
    for i in range(len(nums)):
        for j in range(i, len(nums)):
            if nums[i] != nums[j]:
                flag = False 

    if flag:
        del nums[1:]
        k = 1
        print(nums)
        return k     
    
    for i in range(len(nums)):
        for j in range(i+1, len(nums)):
            if nums[i] == nums[j] and last_num != nums[i]:
                last_num = nums[i]
                cnt_duplicates += 1
                array_1_indexes.append(j)
                break
            elif nums[i] == nums[j] and last_num == nums[i]:
                array_1_indexes.append(j)
                cnt_duplicates += 1

    k = len(nums) - cnt_duplicates
    index_to_delete = -1

    for i in range(1, len(array_1_indexes)):
        if array_1_indexes[i-1] + 1 == array_1_indexes[i]:
            index_to_delete = i

    if index_to_delete != -1 and index_to_delete:
        del array_1_indexes[index_to_delete]

    for i in range(1, len(nums)):
        if nums[i] != nums[i-1]:
            array_2_indexes.append(i)

    for inizio, fine in zip(reversed(array_1_indexes), reversed(array_2_indexes)):
        del nums[inizio:fine]

    print(nums)
    return k

nums_1 = [1, 1, 2]
nums_2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
nums_3 = [1, 1, 1]
nums_4 = [1, 1, 1, 2]

print(removeDuplicates(nums_1), "\n")
print(removeDuplicates(nums_2), "\n")
print(removeDuplicates(nums_3), "\n")
print(removeDuplicates(nums_4))