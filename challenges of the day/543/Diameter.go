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

type binaryTree struct {
	root *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	// diametro "centrale"
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
