class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        max_num = max(candies)
        res = [False]*len(candies)

        for i, c in enumerate(candies):
            if c + extraCandies >= max_num:
                res[i] = True

        return res
