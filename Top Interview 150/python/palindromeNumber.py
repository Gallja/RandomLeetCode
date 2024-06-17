class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        newNum = str(x)
        newNumRev = newNum[::-1]
        for i in range(len(newNum)):
            if newNum[i] != newNumRev[i]:
                return False
        
        return True

solution = Solution()
result = solution.isPalindrome(121)
print(result)