# 数据结构(DataStructure)

---

## **数据结构的存储方式只有两种: 数组(顺序存储) 和链表(链式存储)**

* 递归思想, 自顶向下,由抽象到具体

    * 数据结构,[上层建筑], 数组和链表才是 [结构基础]

    * 线性表(List): 零或N个数据元素的有限序列

        * 数组(ArrayList): 元素内存地址物理相邻
            * 优 : 可随机访问, 可索引快速寻找对应元素, 相对节约存储空间
            * 劣 : 因连续存储, 内存空间需要一次性分配够. 若要扩容, 需新分配一块更大的空间,再把数据全部复制过去, 若想插入 删除,则要搬移后面所有数据 O(N)
            * 时间复杂度: 
                * 查改: O(1)
                * 增删: O(n)

        * 链表(Linklist): 元素内存地址逻辑相邻
            * 优: 因元素不连续,是靠指针指向下一个元素的位置, 所以没有扩容问题. 若知道一元素的前驱和后继,则可操作指针删除此元素或插入新元素
            * 劣: 因存储空间不连续,所以不能通过索引查找元素位置,所以不能随机访问,相对消耗更多的存储空间
            * 时间复杂度:
                * 增删: O(1)
                * 查改: O(n)
            
            * 单链表: 
                *       #! /usr/bin/env python
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
                                    


    
            * 循环链表(Circular Linklist): 单链表的终端节点的指针由null 指向头指针
                * 解决由任意节点即可遍历链表
            * 双向链表(Double Linklist): 单链表 增加一个指向前驱节点的指针 
                * 查找上一节点是O(1), 而不是O(n)


    * 队列 : 先进先出
        * 实现方式
    
           *  数组: 考虑缩容 扩容问题
                * TODO
            * 链表: 不用考虑 缩扩容问题, 但是需要跟多的内存空间存储节点指针
                * TODO

    * 栈 : 先进后出
        * 实现方式
            * 数组: 考虑缩容, 扩容问题
                * TODO
            * 链表: 不用考虑缩扩容问题, 消耗更多的存储空间
                * TODO

    * 图:
        * TODO
    
    * 

---
## 数据结构的基本操作

* 任意数据机构,其操作无非遍历 + 访问(增删查改):
    
    数据结构的种类很多,目的都是为了在不同的应用场景,尽可能高效的增删查改,如何操作

    最高层看, 操作两种形式:线性的and非线性的

    **线性的:for/while 循环(迭代)**

    **非线性的: 递归**

    数组遍历框架: 典型的线性迭代结构

        void traverse(int[] arr){
            for(int i = 0; i < arr.length; i++){
                // 迭代访问arr[i]
            }
        }

    链表遍历框架: 兼具迭代和递归结构:

        //basic Node
        class Node {
            int val;
            Node next;
        }

        void traverse(Node head){
            for (Node p = head; p != null; p = p.next){
                // 迭代访问 p.val
            }
        }

        void tarverse(Node head){
            // 递归访问 p.val
            traverse(head.next);
        }

    二叉树遍历框架,典型非线性递归结构:

        //基本二叉树节点
        class TreeNode{
            int val;
            Treenode left, right;
        }

        void tarverse(TreeNode root){
            traverse(root.left);
            traverse(roo.right);
        }
    
