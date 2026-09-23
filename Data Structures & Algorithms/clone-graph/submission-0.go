/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    oldToNew := make(map[*Node]*Node)

	var dfs func(node *Node)*Node

	dfs = func(node *Node)*Node{
		if node == nil{
			return nil
		}

		if _, found := oldToNew[node]; found{
			return oldToNew[node]
		}
		cpy := &Node{Val: node.Val}
		oldToNew[node] = cpy

		for _, ni := range node.Neighbors{
			cpy.Neighbors = append(cpy.Neighbors, dfs(ni))
		}
		return cpy
	}
	return dfs(node)
}
