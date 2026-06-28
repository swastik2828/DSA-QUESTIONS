class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        low=high=prices[0]
        res=0
        for price in prices:
            if price<low:
                low=price
                high=price
            elif price>high:
                high=price
                res=max(res,high-low)
        return res