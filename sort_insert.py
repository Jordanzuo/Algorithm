#!/usr/bin/pytho3

import time
import random

def insertionSort1(a):
    if len(a) <= 1:
        return;

    for i in range(1, len(a)):
        value = a[i]
        j = i - 1
        while j >= 0:
            if a[j] > value:
                a[j+1] = a[j]
                j = j - 1
            else:
                break
        a[j+1] = value

def insertionSort2(a):
    if len(a) <= 1:
        return;

    for i in range(1, len(a)):
        value = a[i]
        hasSwap = False
        j = i - 1
        while j >= 0:
            if a[j] > value:
                a[j+1] = a[j]
                j = j - 1
                hasSwap = True
            else:
                break
        if hasSwap:
            a[j+1] = value

if __name__ == "__main__":
    a1 = []
    a2 = []
    for i in range(1, 10000):
        num = random.randint(1, 10000)
        a1.append(num)
        a2.append(num)

#    print("Before sort:", a1)
    start = time.time()
    insertionSort1(a1)
    end = time.time()
#    print("After sort:", a1)
    print("used time:", end - start)

#    print("Before sort:", a2)
    start = time.time()
    insertionSort2(a2)
    end = time.time()
#    print("After sort:", a2)
    print("used time:", end - start)
