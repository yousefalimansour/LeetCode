1class Solution(object):
2    def maxSideLength(self, mat, threshold):
3        m, n = len(mat), len(mat[0])
4        prefix = [[0] * (n + 1) for _ in range(m + 1)]
5        max_side = 0
6        
7        for i in range(1, m + 1):
8            for j in range(1, n + 1):
9                prefix[i][j] = (mat[i-1][j-1] 
10                              + prefix[i-1][j] 
11                              + prefix[i][j-1] 
12                              - prefix[i-1][j-1])
13                
14                if i > max_side and j > max_side:
15                    k = max_side + 1
16                    square_sum = (prefix[i][j] 
17                                - prefix[i-k][j] 
18                                - prefix[i][j-k] 
19                                + prefix[i-k][j-k])
20                    
21                    if square_sum <= threshold:
22                        max_side = k
23        
24        return max_side