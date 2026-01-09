1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
10    result, _ := dfs(root)
11    return result
12}
13
14func dfs(node *TreeNode) (*TreeNode, int){
15    if node == nil{
16        return nil, 0
17    }
18    ln , ld := dfs(node.Left)
19    rn , rd := dfs(node.Right)
20
21    if ld > rd {
22        return ln , ld + 1
23    }
24    if rd > ld {
25        return rn , rd + 1
26    }
27
28    return node , ld + 1
29}