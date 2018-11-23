#!/usr/bin/python3
import sys

def binarySort(itemList):
    if len(itemList) == 1: return itemList
    mid = len(itemList) // 2
    left = binarySort(itemList[:mid])
    right = binarySort(itemList[mid:])

    retItemList = []
    leftIndex, leftIndexMax, rightIndex, rightIndexMax = 0, len(left), 0, len(right)
    while leftIndex <= leftIndexMax or rightIndex <= rightIndexMax:
        if leftIndex < leftIndexMax and rightIndex < rightIndexMax:
            if left[leftIndex] < right[rightIndex]:
                retItemList.append(left[leftIndex])
                leftIndex += 1
            else:
                retItemList.append(right[rightIndex])
                rightIndex += 1
        elif leftIndex < leftIndexMax:
            retItemList.extend(left[leftIndex:])
            break
        elif rightIndex < rightIndexMax:
            retItemList.extend(right[rightIndex:])
            break

    return retItemList

if __name__ == "__main__":
    itemMap = map(int, sys.argv[1:])
    itemList = []
    print("Before sort")
    for item in itemMap:
        itemList.append(item)
        print(item)
    itemList = binarySort(itemList)
    print("After sort")
    for item in itemList:
        print(item)
