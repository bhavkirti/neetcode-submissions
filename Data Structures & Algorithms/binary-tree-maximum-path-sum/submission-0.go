/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := []int{root.Val}

	var dfs func(root *TreeNode)int
	dfs = func(root *TreeNode)int{
		if root == nil{
			return 0
		}
		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		leftMax = max(0,leftMax)
		rightMax = max(0, rightMax)

		res[0] = max(res[0], root.Val+ leftMax + rightMax)
		return root.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return res[0]
}

func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}
