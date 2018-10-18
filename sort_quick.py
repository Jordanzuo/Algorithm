#!/usr/bin/python3
import sys

def partition(arr, low, high):
    pivot = arr[high]
    # i points to the first unsorted item
    # j iterates all the unsorted item, and compare with the pivot, if it's less that pivot, move it to the first unsorted position to make it become sorted item. Then i increases to the next unsorted item.
    i = low
    for j in range(low, high):
       if arr[j] < pivot:
            if i != j:
                arr[i], arr[j] = arr[j], arr[i]
            i = i + 1

    # if i doesn't point to pivot, then switch them to make pivot to the right position.(after the last sorted item)
    if i != high:
        arr[i], arr[high] = arr[high], arr[i]

    return i

def quickSort(itemList, low, high):
    if(low < high):
        pi = partition(itemList, low, high)
        quickSort(itemList, low, pi - 1)
        quickSort(itemList, pi + 1, high)

if __name__ == "__main__":
    itemMap = map(int, sys.argv[1:])
    itemList = []
    for item in itemMap:
        itemList.append(item)

    print("Before sort:")
    print("  ".join(map(str, itemList)))
    print("")

    quickSort(itemList, 0, len(itemList) - 1)

    print("")
    print("After sort:")
    print("  ".join(map(str, itemList)))
