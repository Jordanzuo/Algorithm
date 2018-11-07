# If it needs transaction fee, this solution will be better
class Solution1:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if prices == None or len(prices)  <= 1:
            return 0
        prev, i = 0, 0
        profit = 0
        while i < len(prices):
            if i == len(prices) - 1:
                profit = profit + prices[i] - prices[prev] if i > prev else profit
                break
            if prices[i+1] > prices[i]:
                i += 1
                continue
            profit = profit + prices[i] - prices[prev] if i > prev else profit
            i += 1
            prev = i

        return profit

# If there is no fee, this solution will be better
class Solution2:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if prices == None or len(prices)  <= 1:
            return 0
        profit = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i-1]:
                profit += prices[i] - prices[i-1]

        return profit
