# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# This method is intuitive, and it's easy to implement
class Solution1:
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
        # This line is critical to this method
        self.tail.next = None
        while curr:
            curr.next, prev, curr = prev, curr, curr.next
        self.head, self.tail = self.tail, self.head

# This method is intuitive, but is neat.
class Solution2(object):
    def reverseKGroup(self, head, k):
        if not head or not head.next or k < 2: return head

        curr = head
        for _ in range(k):
            if not curr:
                return head
            curr = curr.next

        prev, curr = None, head
        for _ in range(k):
            curr.next, prev, curr = prev, curr, curr.next
        head.next = self.reverseKGroup(curr, k)
        return prev
