package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored")
	}
	node.Value = value
}
