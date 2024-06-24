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

        numFinale = 0
        numProvv = 0
        i = 0

        while i < len(s) - 1:
            if switch(s[i]) < switch(s[i + 1]):
                numProvv = switch(s[i + 1]) - switch(s[i])
                numFinale += numProvv
                i += 1
            else:
                numFinale += switch(s[i])
            
            i += 1
                
        return numFinale
    
solution = Solution()    
print(solution.romanToInt('III')) # 3

print(solution.romanToInt('LVIII')) # 58

print(solution.romanToInt('MCMXCIV')) # 1994