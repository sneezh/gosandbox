package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := 0
	maxGain(root, &res)
	return res
}

func maxGain(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}

	left := maxGain(node.Left, res)
	right := maxGain(node.Right, res)
	currentNodeNotAsRoot := max(max(left, right)+node.Val, node.Val)
	currentNodeAsRoot := max(currentNodeNotAsRoot, left+right+node.Val)
	*res = max(*res, currentNodeAsRoot)
	return currentNodeNotAsRoot
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
