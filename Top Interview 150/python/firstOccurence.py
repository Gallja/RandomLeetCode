def firstOccurence(str_1, str_2):
    cnt = 0
    index = -1

    if not str_1 in str_2 and not str_2 in str_1:
        return index

    if len(str_1) > len(str_2):
        for i in range(len(str_1)):
            for j in range(len(str_2)):
                if str_1[i] == str_2[j] and cnt < 2:
                    cnt += 1
                if str_1[i] == str_2[j] and cnt >= 2:
                    index = i-1
                    break
            if index != -1:
                break
    else:
        for i in range(len(str_2)):
            for j in range(len(str_1)):
                if str_2[i] == str_1[j] and cnt < 2:
                    cnt += 1
                if str_2[i] == str_1[j] and cnt >= 2:
                    index = i-1
                    break
            if index != -1:
                break
    
    return index


str_1 = "sadbutsad" 
str_2 = "sad" 
print(firstOccurence(str_1, str_2))
# output: 0

str_3 = "leetcode" 
str_4 = "leeto"
print(firstOccurence(str_3, str_4))
# output: -1