import heapq

class KthLargest(object):
    def __init__(self, k, nums):
        self.k = k
        self.sortedlist = []
        for num in nums:
            if len(self.sortedlist) < self.k:
                heapq.heappush(self.sortedlist, num)
            elif num > self.sortedlist[0]:
                heapq.heappushpop(self.sortedlist, num)

    def add(self, val):
        if len(self.sortedlist) < self.k:
            heapq.heappush(self.sortedlist, val)
        elif val > self.sortedlist[0]:
            heapq.heappushpop(self.sortedlist, val)
        return self.sortedlist[0]
