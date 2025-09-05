package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = Insert(root.Right, val)
	} else {
		root.Left = Insert(root.Left, val)
	}
	return root
}

// In-order traversal
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Val, " ")
	InOrder(root.Right)
}

func Search(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val < val {
		return Search(root.Right, val)
	}
	return Search(root.Left, val)
}

// PrintTree in a simple visual way
func PrintTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	if root.Right != nil {
		PrintTree(root.Right, prefix+func() string {
			if isLeft {
				return "│   "
			}
			return "    "
		}(), false)
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Println(root.Val)

	if root.Left != nil {
		PrintTree(root.Left, prefix+func() string {
			if isLeft {
				return "    "
			}
			return "│   "
		}(), true)
	}
}

func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(" ")
		fmt.Print(root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Print(" ")
		fmt.Print(root.Val)
	}
}

func main() {
	var root *TreeNode
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		root = Insert(root, v)
	}

	fmt.Println("In-order traversal:")
	InOrder(root)
	fmt.Println("\n\nVisual Tree:")
	PrintTree(root, "", true)

	fmt.Println("\n\nSearch for 4")
	node := Search(root, 4)
	if node != nil {
		fmt.Println("Found:", node.Val)
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("\n\nPre-order traversal:")
	PreOrder(root)

	fmt.Println("\n\nPost-order traversal:")
	PostOrder(root)
}
