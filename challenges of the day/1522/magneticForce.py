from typing import List

class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        position.sort()
        
        def ok(distance):
            cnt, last_position = 1, position[0]

            for pos in position[1:]:
                if pos - last_position >= distance:
                    cnt += 1
                    last_position = pos
                    
                    if cnt == m:
                        return True
                    
            return False
        
        left, right = 1, position[-1] - position[0]
        best_distance = 0
        
        while left <= right:
            mid = left + (right - left) // 2

            if ok(mid):
                best_distance = mid
                left = mid + 1
            else:
                right = mid - 1
        
        return best_distance

solution = Solution()
position = [1, 2, 3, 4, 7]
print(solution.maxDistance(position, 3))

position2 = [5, 4, 3, 2, 1, 1000000000]
print(solution.maxDistance(position2, 2))

position3 = [79, 74, 57, 22]
print(solution.maxDistance(position3, 4))

position4 = [79, 74, 57, 22]
print(solution.maxDistance(position4, 4))