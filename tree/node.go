package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	//fmt.Println(node.value)
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) traverse() {
	if node == nil{
		return
	}
	node.Left.traverse()
	node.Print()
	node.Right.traverse()
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value:value}
}




