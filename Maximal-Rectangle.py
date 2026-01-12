1class Solution(object):
2    def maximalRectangle(self, matrix):
3        if not matrix or not matrix[0]:
4            return 0
5        
6        rows, cols = len(matrix), len(matrix[0])
7        heights = [0] * cols
8        max_area = 0
9        
10        for i in range(rows):
11            for j in range(cols):
12                heights[j] = heights[j] + 1 if matrix[i][j] == '1' else 0
13            
14            max_area = max(max_area, self.largestRectangleArea(heights))
15        
16        return max_area
17    
18    def largestRectangleArea(self, heights):
19        stack = [-1]
20        max_area = 0
21        heights.append(0)
22        
23        for i in range(len(heights)):
24            while heights[i] < heights[stack[-1]]:
25                h = heights[stack.pop()]
26                w = i - stack[-1] - 1
27                max_area = max(max_area, h * w)
28            stack.append(i)
29        
30        heights.pop()
31        return max_area
32        