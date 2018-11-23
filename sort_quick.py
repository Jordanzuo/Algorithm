#!/usr/bin/python3
import sys
import random

def partition(arr, low, high):
    randNum = random.randint(1, 3)
    if randNum == 1:
        pivotIndex = high
    elif randNum == 2:
        pivotIndex = random.randint(low, high)
    else:
        if high - low > 3:
            l, m, r = low, (low + high)//2, high
            _min = min(arr[l], arr[m], arr[r])
            _max = max(arr[l], arr[m], arr[r])
            if arr[l] != _min and arr[l] != _max:
                pivotIndex = l
            if arr[m] != _min and arr[m] != _max:
                pivotIndex = m
            if arr[r] != _min and arr[r] != _max:
                pivotIndex = r
        else:
            pivotIndex = (high + low)//2
    pivot = arr[pivotIndex]

    # i points to the first unsorted item
    # j iterates all the unsorted item, and compare with the pivot, if it's less that pivot, move it to the first unsorted position to make it become sorted item. Then i increases to the next unsorted item.
    i = low
    for j in range(low, high+1):
       if arr[j] < pivot:
            if i != j:
                arr[i], arr[j] = arr[j], arr[i]
                # If the pivot has been switched to another place, then adjust the pivotIndex
                if i == pivotIndex: pivotIndex = j
            i = i + 1

    # if i doesn't point to pivot, then switch them to make pivot to the right position.(after the last sorted item)
    if i != pivotIndex: arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]

    return i

def quickSort(itemList, low, high):
    if(low < high):
        pi = partition(itemList, low, high)
        quickSort(itemList, low, pi - 1)
        quickSort(itemList, pi + 1, high)

if __name__ == "__main__":
    itemMap = map(int, sys.argv[1:])
    itemList = []
    for item in itemMap: itemList.append(item)

    print("Before sort:")
    print("  ".join(map(str, itemList)))
    print("")

    quickSort(itemList, 0, len(itemList) - 1)

    print("")
    print("After sort:")
    print("  ".join(map(str, itemList)))
