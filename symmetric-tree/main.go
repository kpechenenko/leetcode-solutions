// 101. Symmetric Tree
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

// Example 1:
// Input: root = [1,2,2,3,4,4,3]
// Output: true

// Example 2:
// Input: root = [1,2,2,null,3,null,3]
// Output: false

// Constraints:
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100

// Follow up: Could you solve it both recursively and iteratively?
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursiveCheck(root.Right, root.Left)
}

func recursiveCheck(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil) != (right == nil) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return recursiveCheck(left.Left, right.Right) && recursiveCheck(left.Right, right.Left)
}
