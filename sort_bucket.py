#!/usr/bin/python3

def bucketSort(arr):
    #Assume all the data are in the region of [1,10]
    l1 = []
    l2 = []
    l3 = []
    l4 = []
    l5 = []
    l6 = []
    l7 = []
    l8 = []
    l9 = []
    l10 = []


    for item in arr:
        if item == 1:
            l1.append(item)
        elif item == 2:
            l2.append(item)
        elif item == 3:
            l3.append(item)
        elif item == 4:
            l4.append(item)
        elif item == 5:
            l5.append(item)
        elif item == 6:
            l6.append(item)
        elif item == 7:
            l7.append(item)
        elif item == 8:
            l8.append(item)
        elif item == 9:
            l9.append(item)
        elif item == 10:
            l10.append(item)

    result = []
    result.extend(l1)
    result.extend(l2)
    result.extend(l3)
    result.extend(l4)
    result.extend(l5)
    result.extend(l6)
    result.extend(l7)
    result.extend(l8)
    result.extend(l9)
    result.extend(l10)

    return result

if __name__ == "__main__":
    arr = [1, 3, 5, 7, 9, 2, 4, 6, 8, 10, 7, 9, 8, 3, 2, 4]
    print("Before:", arr)
    arr = bucketSort(arr)
    print("After:", arr)
