class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if not nums or len(nums) < 3:
            return []

        nums.sort()
        res = set()
        for i, vi in enumerate(nums[:-2]):
            if i > 0 and vi == nums[i-1]:
                continue
            d = {}
            for vj in nums[i+1:]:
                if vj in d:
                    res.add((vi, -(vi+vj), vj))
                else:
                    d[-(vi+vj)] = 1

        return [item for item in res]
