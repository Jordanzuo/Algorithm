# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True

        t2l = []
        self.tree2list(root, t2l)
        print(t2l)
        for i in range(len(t2l) - 1):
            if t2l[i] >= t2l[i+1]:
                return False
        return True

    def tree2list(self, root, t2l):
        if not root:
            return
        if root.left:
            val = self.tree2list(root.left, t2l)
            if val:
                t2l.append(val)
        t2l.append(root.val)
        if root.right:
            val = self.tree2list(root.right, t2l)
            if val:
                t2l.append(val)
