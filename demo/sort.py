#！/usr/bin/env python
#coding=utf-8

# 快速排序
def quickSort(arr: list):
    if len(arr) <=1:
        return arr

    pivot = arr[0]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    return quickSort(left) + middle + quickSort(right)

def main():
    list1 = [1, 3, 5, 2, 5, 2, 6, 3, 1, 7, 0, 9, 5, 2, 7, 10, 13, 22, 6, 13]
    print("origin arr", list1)
    sorted_list = quickSort(list1)
    print("sorted arr", sorted_list)


if __name__ == "__main__":
    main()
