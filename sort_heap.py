#!/usr/bin/python3
import sys

def heapSort(arr):
    n = len(arr)
    count = 0
    arr.insert(0, 0)

    buildHeap(arr, n)
    k = n
    while k > 1:
        arr[k], arr[1] = arr[1], arr[k]
        k -= 1
        heapify(arr, k, 1)

    return arr

def buildHeap(arr, n):
    for i in range(n//2, 0, -1):
        heapify(arr, n, i)

def heapify(arr, n, i):
    while True:
        maxIndex = i
        if i * 2 <= n and arr[i] < arr[i * 2]: maxIndex = i * 2
        if i * 2 + 1 <= n and arr[maxIndex] < arr[i * 2 + 1]: maxIndex = i * 2 + 1
        if maxIndex == i: break
        arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
        i = maxIndex

if __name__ == "__main__":
    itemMap = map(int, sys.argv[1:])
    itemList = []
    for item in itemMap: itemList.append(item)

    print("Before sort:")
    print("  ".join(map(str, itemList)))
    print("")

    heapSort(itemList)

    print("")
    print("After sort:")
    print("  ".join(map(str, itemList)))

