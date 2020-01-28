package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored")
	}
	node.Value = value
}

func main() {
	var root TreeNode

	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = createNode(2)

	nodes := []TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Right.Left.setValue(4)
	root.Right.Left.print()

	fmt.Println()

	root.print()
	root.setValue(100)

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	root.traverse()

}
