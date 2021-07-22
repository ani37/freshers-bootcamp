package main

import "fmt"

type Tree struct {
	data string
	left *Tree
	right *Tree
}

func (node *Tree)preOrder(){
	if node != nil {
		fmt.Printf(node.data)
		node.left.preOrder()
		node.right.preOrder()
	}

	return
}

func (node *Tree)postOrder(){
	if node != nil {
		node.left.postOrder()
		node.right.postOrder()
		fmt.Printf(node.data)
	}

	return
}

func main(){
	root := Tree{"+", nil, nil}
	root.left = &Tree{"a", nil, nil}
	root.right = &Tree{"-", nil, nil}
	root.right.left = &Tree{"b", nil, nil}
	root.right.right = &Tree{"c", nil, nil}

	fmt.Println("The Preorder of the give tree is")
	root.preOrder()

	fmt.Println("\nThe Postorder of the give tree is")
	root.postOrder()
}