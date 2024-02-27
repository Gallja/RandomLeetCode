package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	// diametro
	diameterThroughRoot := leftHeight + rightHeight

	// diametro sx
	leftDiameter := diameterOfBinaryTree(root.Left)

	// diametro dx
	rightDiameter := diameterOfBinaryTree(root.Right)

	// max diametro
	return max(diameterThroughRoot, max(leftDiameter, rightDiameter))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// altezza massima tra sottoalbero sx e dx
	return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(root))
}
