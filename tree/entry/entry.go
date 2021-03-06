package main

import (
	"go/tree"
	"fmt"
	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct{
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
    if myNode == nil || myNode.node == nil{
    	return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse()  {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

/*
     3
   /   \
  0     5
   \   /
    2 0
*/

func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
	testSparse()

}