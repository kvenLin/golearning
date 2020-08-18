package main

import (
	"fmt"
	"imooc.com/louye/learngo/tree"
)

type myTeeNode struct {
	node *tree.Node
}

func (myNode *myTeeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTeeNode{myNode.node.Left}
	left.postOrder()
	right := myTeeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Left.Right = tree.CreateNode(2)
	root.Traverse()
	myNode := myTeeNode{&root}
	myNode.postOrder()

}
