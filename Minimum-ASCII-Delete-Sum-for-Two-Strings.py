1class Solution(object):
2    def minimumDeleteSum(self, s1, s2):
3        """
4        :type s1: str
5        :type s2: str
6        :rtype: int
7        """
8        n, m = len(s1), len(s2)
9        dp = [[0] * (m + 1) for _ in range(n + 1)]
10        
11        for j in range(1, m + 1):
12            dp[0][j] = dp[0][j-1] + ord(s2[j-1])
13        
14        for i in range(1, n + 1):
15            dp[i][0] = dp[i-1][0] + ord(s1[i-1])
16        
17        for i in range(1, n + 1):
18            for j in range(1, m + 1):
19                if s1[i-1] == s2[j-1]:
20                    dp[i][j] = dp[i-1][j-1]
21                else:
22                    dp[i][j] = min(
23                        dp[i-1][j] + ord(s1[i-1]),  
24                        dp[i][j-1] + ord(s2[j-1])   
25                    )
26        
27        return dp[n][m]