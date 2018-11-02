class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        if not nums or len(nums) < 2:
            return None

        d = {}
        for i, num in enumerate(nums):
            prev = d.get(target-num)
            if prev != None:
                return [prev, i]
            if d.get(num) == None:
                d[num] = i
        return None
