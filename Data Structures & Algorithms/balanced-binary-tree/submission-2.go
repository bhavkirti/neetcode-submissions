/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Result struct{
	balanced bool
	height int
}

func isBalanced(root *TreeNode) bool {
    return dfs(root).balanced
}

func dfs(root *TreeNode) Result {
	if root == nil {
		return Result{true, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	balanced := left.balanced && right.balanced && abs(left.height - right.height) <= 1
	return Result{balanced, 1+max(left.height, right.height)}
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}

func abs(a int) int{
	if a > 0 {
		return a
	}
	return -a
}
