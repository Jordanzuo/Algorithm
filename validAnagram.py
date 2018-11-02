class Solution:
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if not s and not t:
            return True
        if not s or not t:
            return False
        if len(s) != len(t):
            return False

        s_dict, t_dict = {}, {}
        for item in s:
            s_dict[item] = s_dict.get(item, 0) + 1
        for item in t:
            t_dict[item] = t_dict.get(item, 0) + 1
        return s_dict == t_dict
