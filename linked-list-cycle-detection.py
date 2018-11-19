# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def detectCycle(self, head):
        if not head:
            return None
        slow = fast = head
        while slow and fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow == fast:
                break
        if not fast or not fast.next:
            return None

        slow = head
        while slow is not fast:
            slow, fast = slow.next, fast.next
        return slow
