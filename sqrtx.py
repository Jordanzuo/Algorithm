#!/usr/bin/python3

class Solution:
	def mySqrt(self, x):
		if x == 0 or x == 1: return x
		left, right = 1, x
		while left <= right:
			mid = (left + right) / 2
			print(left, mid, right)
			if mid == x / mid:
				return mid
			elif mid < x / mid:
				left = mid
			else:
				right = mid
			if abs(left - right) <= 0.05:
				break
		print(left, mid, right)
		return int(right)

if __name__ == "__main__":
	solution = Solution()
	print(solution.mySqrt(2147395599))
