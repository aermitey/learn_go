package main

import "fmt"

func main() {
	root := buildDLink()
	root = insertNode(root, &LinkNode{data: -1})
	root = insertNode(root, &LinkNode{data: 0})
	root = insertNode(root, &LinkNode{data: 1})
	root = insertNode(root, &LinkNode{data: 3})
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 10})
	root = insertNode(root, &LinkNode{data: 15})
	rangeLink(root)
	fmt.Println("删除")
	root = deleteNode(root, 1)
	root = deleteNode(root, 10)
	root = deleteNode(root, 15)
	rangeLink(root)
}

type LinkNode struct {
	data     int
	next     *LinkNode
	previous *LinkNode
}

func buildDLink() *LinkNode {
	n1 := &LinkNode{data: 1}
	n2 := &LinkNode{data: 5}
	n3 := &LinkNode{data: 10}

	n1.next = n2
	n2.previous = n1

	n2.next = n3
	n3.previous = n2

	return n1
}

func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	for {
		fmt.Println(tmpNode.data)
		if tmpNode.previous == nil {
			break
		}
		tmpNode = tmpNode.previous
	}
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root
	if root == nil {
		return newNode
	}
	if newNode.data < root.data {
		newNode.next = tmpNode
		tmpNode.previous = newNode
		return newNode
	}
	for {
		if tmpNode.next == nil {
			//已经到头了，追加节点
			tmpNode.next = newNode
			newNode.previous = tmpNode
			return root
		} else if tmpNode.next.data > newNode.data {
			//在此插入新节点
			newNode.next = tmpNode.next
			newNode.previous = tmpNode
			tmpNode.next.previous = newNode
			tmpNode.next = newNode
			return root
		}
		tmpNode = tmpNode.next
	}
}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	if root == nil {
		return nil
	}
	if root.data == data {
		if root.next == nil {
			return nil
		}
		rightNode := tmpNode.next
		tmpNode.next = nil
		//清理rightNode的previous
		rightNode.previous = nil
		return rightNode
	}
	for {
		if tmpNode.next == nil {
			return root
		} //else if tmpNode.next.data == data {
		//	rightNode := tmpNode.next
		//	tmpNode.next = rightNode.next
		//	rightNode.next.previous = tmpNode
		//	return root
		//}
		rightNode := tmpNode.next
		if rightNode.data == data {
			if rightNode.next != nil {
				tmpNode.next = rightNode.next
				rightNode.next.previous = tmpNode
				//清理右游标的关联关系
				rightNode.next = nil
				rightNode.previous = nil
				return root
			} else {
				rightNode.previous = nil
				tmpNode.next = nil
				return root
			}
		}
		tmpNode = tmpNode.next
	}
}
