1func maxDotProduct(nums1 []int, nums2 []int) int {
2    n, m := len(nums1), len(nums2)
3    
4    // Quick check for all negative vs all positive
5    allNeg1, allPos1 := true, true
6    allNeg2, allPos2 := true, true
7    max1, min1 := nums1[0], nums1[0]
8    max2, min2 := nums2[0], nums2[0]
9    
10    for _, v := range nums1 {
11        if v >= 0 {
12            allNeg1 = false
13        }
14        if v <= 0 {
15            allPos1 = false
16        }
17        if v > max1 {
18            max1 = v
19        }
20        if v < min1 {
21            min1 = v
22        }
23    }
24    
25    for _, v := range nums2 {
26        if v >= 0 {
27            allNeg2 = false
28        }
29        if v <= 0 {
30            allPos2 = false
31        }
32        if v > max2 {
33            max2 = v
34        }
35        if v < min2 {
36            min2 = v
37        }
38    }
39    
40    // Edge case optimization
41    if (allNeg1 && allPos2) || (allPos1 && allNeg2) {
42        if allNeg1 && allPos2 {
43            return max1 * min2
44        }
45        return min1 * max2
46    }
47    
48    // In-place DP with single array
49    dp := make([]int, m)
50    dp[0] = nums1[0] * nums2[0]
51    
52    for j := 1; j < m; j++ {
53        dp[j] = max(dp[j-1], nums1[0]*nums2[j])
54    }
55    
56    for i := 1; i < n; i++ {
57        prev := dp[0]
58        dp[0] = max(dp[0], nums1[i]*nums2[0])
59        
60        for j := 1; j < m; j++ {
61            temp := dp[j]
62            product := nums1[i] * nums2[j]
63            dp[j] = max(
64                max(product, prev+product),
65                max(dp[j], dp[j-1]),
66            )
67            prev = temp
68        }
69    }
70    
71    return dp[m-1]
72}
73
74func max(a, b int) int {
75    if a > b {
76        return a
77    }
78    return b
79}