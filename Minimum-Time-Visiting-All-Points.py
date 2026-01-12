1class Solution(object):
2    def minTimeToVisitAllPoints(self, points):
3        return sum(max(abs(points[i+1][0] - points[i][0]), 
4                      abs(points[i+1][1] - points[i][1]))
5                  for i in range(len(points) - 1))
6        