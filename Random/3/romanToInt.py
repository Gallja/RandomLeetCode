class Solution:
    def romanToInt(self, s: str) -> int:
        def switch(car) -> int:
            if car == 'I':
                return 1
            elif car == 'V':
                return 5
            elif car == 'X':
                return 10
            elif car == 'L':
                return 50
            elif car == 'C':
                return 100
            elif car == 'D':
                return 500
            elif car == 'M':
                return 1000

        return 0
    
solution = Solution()    
print(solution.romanToInt('III'))