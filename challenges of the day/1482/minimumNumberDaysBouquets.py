from typing import List
import sys

class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        l = 1
        r = sys.maxsize
        numDays = -1

        while l <= r:
            mid = l + (r - l) // 2
            consecutive_length, bouquets = 0, 0

            for bloom in bloomDay:
                if bloom <= mid:
                    consecutive_length += 1
                    if consecutive_length >= k:
                        consecutive_length = 0
                        bouquets += 1
                else:
                    consecutive_length = 0

            if bouquets >= m:
                numDays = mid
                r = mid - 1
            else:
                l = mid + 1

        return numDays
    

solution = Solution()
bloomDay = [1, 10, 3, 10, 2]
m = 3 # bouquets to make
k = 1 # adjacent flowers
print(solution.minDays(bloomDay, m, k))

solution2 = Solution()
bloomDay2 = [1, 10, 3, 10, 2]
print(solution2.minDays(bloomDay2, 2, 3))

solution3 = Solution()
bloomDay3 = [7, 7, 7, 7, 12, 7, 7]
print(solution3.minDays(bloomDay3, 2, 3))