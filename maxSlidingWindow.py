#/usr/bin/python3

class Solution(object):
    def maxSlidingWindow(self, nums, k):
        if len(nums) == 0:
            return []
        elif len(nums) < k:
            return []
        elif k == 1:
            return nums
        elif len(nums) == k:
            maxValue = nums[0]
            for num in nums:
                if num > maxValue:
                    maxValue = num
            return [maxValue]

        retList = []
        tmpList = []
        maxValue = nums[0]
        maxIndex = 0
        for index, num in enumerate(nums[1:k]):
            if num >= maxValue:
                maxValue = num
                maxIndex = index
        tmpList.extend(nums[maxIndex:k])
        retList.append(maxValue)

        nums = nums[k:]
        if len(nums) == 0:
            return retList
        for num in nums:
            if num > maxValue:
                maxValue = num
                tmpList.clear()
                tmpList.append(num)
                retList.append(maxValue)
            else:
                tmpList.append(num)
                if len(tmpList) <= k:
                    retList.append(maxValue)
                    continue
                first = tmpList.pop(0)
                if first != maxValue:
                    retList.append(maxValue)
                else:
                    maxValue = tmpList[0]
                    maxIndex = 0
                    for index, num in enumerate(tmpList[1:]):
                        if num >= maxValue:
                            maxValue = num
                            maxIndex = index
                    tmpList = tmpList[maxIndex:]
                    retList.append(maxValue)

        return retList

if __name__ == "__main__":
    nums = [1, 3, -1, -3, 5, 3, 6, 7]
    k = 3
    s = Solution()
    print(s.maxSlidingWindow(nums, k))
