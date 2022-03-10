package main

import "fmt"

type Node struct {
	data  int
	root  *Node
	left  *Node
	right *Node
}

func main() {
	root := buildTree()
	root = insertNode(root, &Node{data: 43})
	root = insertNode(root, &Node{data: 28})
	root = deleteNode(root, 51)
	root = deleteNode(root, 43)
	fmt.Println("done")
}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}

	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1
	return n1
}

func insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			root.left = newNode
			newNode.root = root
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.data > root.data {
		if root.right == nil {
			root.right = newNode
			newNode.root = root
		} else {
			insertNode(root.right, newNode)
		}
	}
	return root
}

//中间状态，忽略
//func deleteNodeLeaf(root *Node, data int) *Node {
//	leftRoot := root
//	if leftRoot.data == data && leftRoot.left == nil && leftRoot.right == nil {
//		leftRoot = leftRoot.root
//		right := root
//		if leftRoot.left == right {
//			//删除左边叶子结点
//			leftRoot.left = nil
//			right.root = nil
//			return leftRoot
//		} else {
//			//删除右边叶子结点
//			leftRoot.left = nil
//			right.root = nil
//			return leftRoot
//		}
//	}
//	return root
//}zh中间状态，忽略

func deleteNode(root *Node, data int) *Node {
	if data < root.data {
		deleteNode(root.left, data)
	} else if data > root.data {
		deleteNode(root.right, data)
	} else {
		//现在root指向要删除的节点
		leftNextGen := findNextGenFromLeft(root.left)
		rightNextGen := findNextGenFromRight(root.right)
		if leftNextGen == nil && rightNextGen == nil {
			//现在要删除的是叶子结点，即最底部的节点
			top := root.root
			down := root
			if top.left == down {
				//表示是左子树
				top.left = nil
				down.root = nil
				return nil
			} else {
				//表示是又子树
				top.right = nil
				down.root = nil
				return nil
			}
		} else if leftNextGen != nil {
			root.data = leftNextGen.data
			deleteNode(leftNextGen, leftNextGen.data)
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data)
		}
	}
	return root
}

func findNextGenFromLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

func findNextGenFromRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}
