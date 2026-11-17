/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var helper func(node *TreeNode, level int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level == len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
		helper(node.Left, level+1)
		helper(node.Right, level+1)
		return
	}

    helper(root, 0)
	return res
}
