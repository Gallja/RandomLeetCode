from typing import List

class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        numDays = 0

        bloomDay.sort()

        if bloomDay[0] < k:
            return -1

        for bouquet in bloomDay:
            while m > 0:
                if bouquet > k:
                    break
                else:
                    numDays += k
                m -= 1

        return numDays
    

solution = Solution()
bloomDay = [1, 10, 3, 10, 2]
m = 3
k = 1
print(solution.minDays(bloomDay, m, k))

solution2 = Solution()
bloomDay2 = [1,10,3,10,2]
print(solution2.minDays(bloomDay2, 2, 3))

solution3 = Solution()
bloomDay3 = [7,7,7,7,12,7,7]
print(solution3.minDays(bloomDay3, 2, 3))