from typing import List

class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        if m >= len(position):
            return 0

        if m == 2:
            return (1 - position[len(position)-1]) * -1
        
        ris = []

        for i, n in enumerate(position):
            if n % m == 0:
                ris.append(i)

        if m == 3:
            return int(ris[0]) + 1
        else:    
            return int(ris[0])

solution = Solution()
position = [1, 2, 3, 4, 7]
print(solution.maxDistance(position, 3))

position2 = [5, 4, 3, 2, 1, 1000000000]
print(solution.maxDistance(position2, 2))

position3 = [79, 74, 57, 22]
print(solution.maxDistance(position3, 4))