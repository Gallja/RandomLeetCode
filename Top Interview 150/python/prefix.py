def longestCommonPrefix(strs):
    if len(strs) == 0:
        return ""

    if len(strs) == 1:
        return strs[0]

    min = strs[0]
    
    for i in range(len(strs)):
        if len(strs[i]) < len(min):
            min = strs[i]
    
    prefix = ""

    for j in range(len(min)):
        for i in range(len(strs)):
            if strs[i][j] != min[j]:
                return prefix
        prefix += min[j]

    return prefix
            

strs_1 = ["flower","flow","flight"]
print(longestCommonPrefix(strs_1))

strs_2 = ["a"]
print(longestCommonPrefix(strs_2))

strs_3 = ["ab", "a"]
print(longestCommonPrefix(strs_3))

strs_4 = ["cir","car"]
print(longestCommonPrefix(strs_4))

strs_5 = ["flower","flower","flower","flower"]
print(longestCommonPrefix(strs_5))