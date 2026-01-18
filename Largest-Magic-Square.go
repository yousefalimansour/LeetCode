1func largestMagicSquare(grid [][]int) int {
2    m, n := len(grid), len(grid[0])
3    rowSum := make([][]int, m)
4    colSum := make([][]int, n)
5    
6    for i := 0; i < m; i++ {
7        rowSum[i] = make([]int, n+1)
8        for j := 0; j < n; j++ {
9            rowSum[i][j+1] = rowSum[i][j] + grid[i][j]
10        }
11    }
12    
13    for j := 0; j < n; j++ {
14        colSum[j] = make([]int, m+1)
15        for i := 0; i < m; i++ {
16            colSum[j][i+1] = colSum[j][i] + grid[i][j]
17        }
18    }
19    
20    for k := min(m, n); k > 1; k-- {
21        for i := 0; i <= m-k; i++ {
22            for j := 0; j <= n-k; j++ {
23                val := rowSum[i][j+k] - rowSum[i][j]
24                valid := true
25                
26                for r := i; r < i+k && valid; r++ {
27                    if rowSum[r][j+k]-rowSum[r][j] != val {
28                        valid = false
29                    }
30                }
31                
32                for c := j; c < j+k && valid; c++ {
33                    if colSum[c][i+k]-colSum[c][i] != val {
34                        valid = false
35                    }
36                }
37                
38                if valid {
39                    d1, d2 := 0, 0
40                    for d := 0; d < k; d++ {
41                        d1 += grid[i+d][j+d]
42                        d2 += grid[i+d][j+k-1-d]
43                    }
44                    if d1 == val && d2 == val {
45                        return k
46                    }
47                }
48            }
49        }
50    }
51    
52    return 1
53}
54
55func min(a, b int) int {
56    if a < b {
57        return a
58    }
59    return b
60}