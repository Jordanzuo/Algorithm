package maxRegionSum

import (
	"math"
)

func GetMaxRegionSum1(list []int) (lowerIndex int, upperIndex int) {
	// Check if all the values are non-positive numbers?
	// If true, just choose the largest number.
	maxValue := math.MinInt
	maxIndex := 0
	firstPositiveIndex := -1
	lastPositiveIndex := -1
	for i, v := range list {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}

		if v > 0 {
			if firstPositiveIndex == -1 {
				firstPositiveIndex = i
			}
			lastPositiveIndex = i
		}
	}

	if maxValue <= 0 {
		lowerIndex = maxIndex
		upperIndex = maxIndex
		return
	}

	// Start from the first positive number until the last positive number.
	// Because the non-positive numbers before the first positive number and after the last positive number doesn't affect the result.
	maxSum := math.MinInt
	for i := firstPositiveIndex; i <= lastPositiveIndex; i++ {
		sum := list[i]
		if sum > maxSum {
			maxSum = sum
			lowerIndex, upperIndex = i, i
		}
		for j := i + 1; j <= lastPositiveIndex; j++ {
			sum += list[j]
			if sum > maxSum {
				maxSum = sum
				lowerIndex, upperIndex = i, j
			}
		}
	}

	return
}

func GetMaxRegionSum2(list []int) (lowerIndex int, upperIndex int) {
	// Check if all the values are non-positive numbers?
	// If true, just choose the largest number.
	maxValue := math.MinInt
	maxIndex := 0
	firstPositiveIndex := -1
	lastPositiveIndex := -1
	for i, v := range list {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}

		if v > 0 {
			if firstPositiveIndex == -1 {
				firstPositiveIndex = i
			}
			lastPositiveIndex = i
		}
	}

	if maxValue <= 0 {
		lowerIndex = maxIndex
		upperIndex = maxIndex
		return
	}

	// Start from the first positive number until the last positive number.
	// Because the non-positive numbers before the first positive number and after the last positive number doesn't affect the result.

	// Split list into regions by a very big negative number which makes the sum from the beginning until now as negative number
	regionList := make([]Region, 0)
	sum := 0
	regionLeftIndex := firstPositiveIndex
	for i := firstPositiveIndex; i <= lastPositiveIndex; {
		sum += list[i]
		if sum >= 0 {
			i++
			continue
		}

		// Add data before i as a region
		sum = 0
		regionList = append(regionList, newRegion(regionLeftIndex, i-1))

		// Skip the following negative numbers
		j := i + 1
		for ; i <= lastPositiveIndex; j++ {
			if list[j] > 0 {
				break
			}
		}

		regionLeftIndex = j
		i = j
	}

	// Add the left numbers as a region
	regionList = append(regionList, newRegion(regionLeftIndex, lastPositiveIndex))

	// If there is no valid region, make the default region
	if len(regionList) == 0 {
		regionList = append(regionList, newRegion(firstPositiveIndex, lastPositiveIndex))
	}

	// Calc max sum for each region
	for i := 0; i < len(regionList); i++ {
		// Get the right most index from left to right
		maxSum := math.MinInt
		sum := 0
		for j := regionList[i].LeftIndex; j <= regionList[i].RightIndex; j++ {
			sum += list[j]
			if sum > maxSum {
				maxSum = sum
				regionList[i].UpperIndex = j
			}
		}

		// Get the left most index from right to left
		maxSum = math.MinInt
		sum = 0
		for j := regionList[i].UpperIndex; j >= regionList[i].LeftIndex; j-- {
			sum += list[j]
			if sum > maxSum {
				maxSum = sum
				regionList[i].LowerIndex = j
			}
		}

		regionList[i].MaxSum = maxSum
	}

	maxSum := math.MinInt
	for _, region := range regionList {
		if region.MaxSum > maxSum {
			maxSum = region.MaxSum
			lowerIndex, upperIndex = region.LowerIndex, region.UpperIndex
		}
	}

	return
}

type Region struct {
	LeftIndex  int
	RightIndex int
	LowerIndex int
	UpperIndex int
	MaxSum     int
}

func newRegion(leftIndex, rightIndex int) Region {
	return Region{
		LeftIndex:  leftIndex,
		RightIndex: rightIndex,
	}
}
