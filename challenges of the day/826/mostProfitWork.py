class Solution(object):
    def maxProfitAssignment(self, difficulty, profit, worker):
        worker.sort()

        coppie = list(zip(difficulty, profit))
        coppie.sort(key=lambda x : x[0])

        maxProfit = 0
        # print(coppie)

        for i in range(len(worker)):
            for j in range(len(coppie)):
                if coppie[j][0] <= worker[i]:
                    maxProfit += coppie[j][1]

        return maxProfit


def generaCoppie(difficulty, profit):
    coppie = []

    for i in range(len(difficulty)):
        coppie.append(Coppia(difficulty[i], profit[i]))

    return coppie

difficulty = [2, 4, 6, 8, 10]
profit = [10, 20, 30, 40, 50]
worker = [4, 5, 6, 7]

solution = Solution()
res1 = solution.maxProfitAssignment(difficulty, profit, worker)
print(res1)