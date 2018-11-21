#/usr/bin/python3

class Solution(object):
    def maxSlidingWindow(self, nums, k):
        length = len(nums)
        if length == 0: return []
        if k == 1: return nums
        if length < k: return []

        # calc the maxNum and corresponding index for the first k nums
        retList = []
        maxNum, maxIndex = nums[0], 0

        # iterate the list, go one step each time
        lower, upper = 0, 0
        for index, num in enumerate(nums[0:]):
            if index == 0: continue
            if index >= k: lower += 1
            upper = index

            if upper - maxIndex < k:
                if num >= maxNum:
                    maxNum = num
                    maxIndex = index
            else:
                if num >= maxNum:
                    maxNum = num
                    maxIndex = index
                else:
                    # recalc the maxNum
                    maxNum = nums[lower]
                    maxIndex = lower
                    for _index, _num in enumerate(nums[lower + 1 : upper+1]):
                        if _num > maxNum:
                            maxNum = _num
                            maxIndex = lower + 1 + _index

            # add the maxNum into the retList
            if index < k - 1: continue
            retList.append(maxNum)

        return retList
