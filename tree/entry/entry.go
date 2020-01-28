package main

import (
	"fmt"

	"haozhexu.com/gofun/learngo/tree"
)

type myTreeNode struct {
	*tree.Node //embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In order traversal: ")
	fmt.Println("root.Traverse(): ")
	root.Traverse()
	fmt.Println("root.Node.Traverse(): ")
	root.Node.Traverse()
	fmt.Println()

	fmt.Println("My own post-order traversal: ")
	root.postOrder()
	fmt.Println()

}
