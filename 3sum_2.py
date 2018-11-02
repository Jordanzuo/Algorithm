class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        if not nums or len(nums) < 3:
            return res

        nums.sort()
        for i, num in enumerate(nums):
            if i > 0 and num == nums[i-1]:
                continue
            l, r = i+1, len(nums)-1
            while l < r:
                if num + nums[l] + nums[r] < 0:
                    l += 1
                elif num + nums[l] + nums[r] > 0:
                    r -= 1
                else:
                    res.append(([num, nums[l], nums[r]]))
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1
                    l, r = l+1, r-1
        return res
