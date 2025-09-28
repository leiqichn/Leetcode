class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        ans = ''
        n = len(word1)
        m = len(word2)

        for i=0; i < n or i <m ;i++: # 错误 python中for 使用range
            if i < n:
                ans += word1[i]
            if i < m:
                ans += word2[i]

        return ans