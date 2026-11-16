/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if subRoot == nil {
		return true
	}

	if isSameTree(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}

}

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val {
		return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
	} else {
		return false
	}
}
