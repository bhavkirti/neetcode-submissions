/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var arr []int
	
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {
			return
		}
		arr = append(arr, root.Val)
	
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	sort.Ints(arr)
	return arr[k-1]
}
