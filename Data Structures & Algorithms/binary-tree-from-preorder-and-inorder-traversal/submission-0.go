/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    pre, in := 0,0
	var dfs func(int) *TreeNode
	dfs = func(limit int) *TreeNode{
		if pre >= len(preorder){
			return nil
		}
		if limit == inorder[in]{
			in++
			return nil
		}
		root := &TreeNode{Val: preorder[pre]}
		pre++

		root.Left = dfs(root.Val)
		root.Right = dfs(limit)
		return root
	}
	return dfs(math.MaxInt)
}
