#! /usr/bin/env python
# coding=utf-8

__author__ = "TF00"


class Node():
    # p: pointer data: data
    def __init__(self, data=None):
        self.data = data
        self.next = None

class SLinkList():
    def __init__(self):
        self.head = None

    def printlist(self):
        temp = self.head    
        while temp is not None:
            print(temp.data)
            temp = temp.next

if __name__ == "__main__":
    # Start with the empty list
    list1 = SLinkList()
    list1.head = Node("1")
    sec = Node("2")
    thr = Node("3")
    list1.head.next = sec
    sec.next = thr
    list1.printlist()
            


    