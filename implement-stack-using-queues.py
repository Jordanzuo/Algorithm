from collections import deque
class MyStack(object):
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.queue1 = deque()
        self.queue2 = deque()

    def push(self, x):
        """
        Push element x onto stack.
        :type x: int
        :rtype: void
        """
        if self.queue1:
            self.queue1.append(x)
        else:
            self.queue2.append(x)

    def pop(self):
        """
        Removes the element on top of the stack and returns that element.
        :rtype: int
        """
        if len(self.queue1) > 0:
            while len(self.queue1) > 1:
                self.queue2.append(self.queue1.popleft())
            return self.queue1.popleft()
        if len(self.queue2) > 0:
            while len(self.queue2) > 1:
                self.queue1.append(self.queue2.popleft())
            return self.queue2.popleft()
        return None

    def top(self):
        """
        Get the top element.
        :rtype: int
        """
        if len(self.queue1) > 0:
            while len(self.queue1) > 1:
                self.queue2.append(self.queue1.popleft())
            ret = self.queue1.popleft()
            self.queue2.append(ret)
            return ret
        if len(self.queue2) > 0:
            while len(self.queue2) > 1:
                self.queue1.append(self.queue2.popleft())
            ret = self.queue2.popleft()
            self.queue1.append(ret)
            return ret
        return None

    def empty(self):
        """
        Returns whether the stack is empty.
        :rtype: bool
        """
        return not self.queue1 and not self.queue2
