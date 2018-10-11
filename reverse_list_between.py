# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseBetween(self, head, m, n):
        if not head or not head.next or m < 1 or m > n:
            return head
        if m == n:
            return head

        part1_begin = part1_end = None
        for i in range(1, m):
            if i == 1:
                part1_begin = head
            part1_end, head = head, head.next

        prev, curr = None, head
        for i in range(m, n+1):
            curr.next, prev, curr = prev, curr, curr.next
        head.next = curr

        if part1_begin == None:
            return prev
        else:
            part1_end.next = prev
            return part1_begin
