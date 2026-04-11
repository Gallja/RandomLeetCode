def mergeTwoList(self, list1, list2):
    min_list = []
    max_list = []
    
    if list1.__len__() > list2.__len__():
        min_list = list2
        max_list = list1
    else:
        min_list = list1
        max_list = list2
    
    index_min = 0
    index_max = 0

    while index_max < len(max_list):
        if max_list[index_max] <= min_list[index_min]:
            self.append(max_list[index_max])
            index_max += 1
        else:
            self.append(min_list[index_min])
            index_min += 1

    for elem in max_list[index_max:]:
        self.append(elem)
        
    for elem in min_list[index_min:]:
        self.append(elem)

    return self

list1 = [1, 2, 4]
list2 = [1, 3, 4]
self = []

print(mergeTwoList(self, list1, list2))