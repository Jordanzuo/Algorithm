# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseKGroup(self, head, k):
        if not head or k == 1:
            return head
        groupList = []
        i = 1
        curr = head
        while curr:
            if i % k == 1:
                groupHead = curr
            groupTail = curr
            if i % k == 0:
                groupList.append(ListItem(groupHead, groupTail, True))
            i = i + 1
            curr = curr.next

        if (i - 1) % k != 0:
            groupList.append(ListItem(groupHead, groupTail, False))

        for item in groupList:
            item.reverse()
        for i in range(len(groupList) - 1):
            groupList[i].tail.next = groupList[i+1].head

        return groupList[0].head

class ListItem:
    def __init__(self, head, tail, ifReverse):
        self.head = head
        self.tail = tail
        self.ifReverse = ifReverse

    def reverse(self):
        if self.ifReverse == False:
            return
        prev, curr = None, self.head
        self.tail.next = None
        while curr:
            curr.next, prev, curr = prev, curr, curr.next
        self.head, self.tail = self.tail, self.head
