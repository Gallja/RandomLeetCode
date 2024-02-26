package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type binaryTree struct {
	root *TreeNode
}

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pContent := recursiveDFSinorder(p)
	qContent := recursiveDFSinorder(q)

	if len(qContent) != len(pContent) {
		return false
	}

	for i := 0; i < len(qContent); i++ {
		if pContent[i] != qContent[i] {
			return false
		}
	}

	return true
}

func newNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func recursiveDFSinorder(root *TreeNode) []int {
	var result []int

	if root != nil {
		// Visita il sottoalbero sinistro
		result = append(result, recursiveDFSinorder(root.Left)...)

		// Aggiungi il valore del nodo corrente
		result = append(result, root.Val)

		// Visita il sottoalbero destro
		result = append(result, recursiveDFSinorder(root.Right)...)
	}

	return result
}

func main() {
	tree1 := binaryTree{}
	tree2 := binaryTree{}

	tree1.root = newNode(1)
	tree1.root.Left = newNode(2)
	tree1.root.Right = newNode(3)

	tree2.root = newNode(1)
	tree2.root.Left = newNode(2)
	tree2.root.Right = newNode(3)

	fmt.Println(recursiveDFSinorder(tree1.root))
	fmt.Println(recursiveDFSinorder(tree2.root))

	fmt.Println(isSameTree(tree1.root, tree2.root))
}
