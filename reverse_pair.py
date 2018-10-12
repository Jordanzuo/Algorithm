# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def swapPairs(self, head):
        first = head
        head = first.next if first and first.next else first
        while first and first.next:
            second = first.next
            third = first.next.next

            second.next = first
            first.next = third.next if third and third.next else third
            first = third

        return head
