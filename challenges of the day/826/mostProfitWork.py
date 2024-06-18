class Solution(object):
    def maxProfitAssignment(self, difficulty, profit, worker):
        worker.sort()

        coppie = list(zip(difficulty, profit))
        coppie.sort(key=lambda x : x[0])

        maxProfit = 0

        for i in range(len(worker)):
            newProfit = False
            profitTmp = 0
            profitPrec = 0
            for j in range(len(coppie)):
                if coppie[j][0] <= worker[i]:
                    if newProfit:
                        if coppie[j][1] > profitPrec:
                            profitPrec = coppie[j][1]
                            profitTmp = 0
                            profitTmp += coppie[j][1]
                    else:
                        profitPrec = coppie[j][1]
                        profitTmp += coppie[j][1]
                        newProfit = True
                else:
                    break
            maxProfit += profitTmp        
            
        return maxProfit

difficulty = [2, 4, 6, 8, 10]
profit = [10, 20, 30, 40, 50]
worker = [4, 5, 6, 7]

solution = Solution()
print(solution.maxProfitAssignment(difficulty, profit, worker))


diff2 = [85,47,57]
prof2 = [24,66,99]
work2 = [40,25,25]

solution2 = Solution()
print(solution2.maxProfitAssignment(diff2, prof2, work2))

diff3 = [68,35,52,47,86]
prof3 = [67,17,1,81,3]
work3 = [92,10,85,84,82]

solution3 = Solution()
print(solution3.maxProfitAssignment(diff3, prof3, work3))