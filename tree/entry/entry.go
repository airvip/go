package main

func main() {
	var root tree.TreeNode


	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.traverse()

}