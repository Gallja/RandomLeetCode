class Solution(object):
    def maxProfitAssignment(self, difficulty, profit, worker):
        """
        :type difficulty: List[int]
        :type profit: List[int]
        :type worker: List[int]
        :rtype: int
        """
        

class Coppia:
    def __init__(self, difficulty, profit):
        self.difficulty = difficulty
        self.profit = profit


difficulty = [2, 4, 6, 8, 10]
profit = [10, 20, 30, 40, 50]
worker = [4, 5, 6, 7]

solution = Solution()
result = solution.maxProfitAssignment(difficulty, profit, worker)