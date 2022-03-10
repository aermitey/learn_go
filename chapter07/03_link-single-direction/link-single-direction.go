package main

import (
	"fmt"
)

type LinkNode struct {
	Data int
	Next *LinkNode
}

func InsertLink(root *LinkNode, node *LinkNode) *LinkNode {
	tmpNode := root
	if root == nil {
		return node
	}
	if root.Data > node.Data {
		node.Next = root
		return node
	}
	for {
		if tmpNode.Next == nil {
			tmpNode.Next = node
			break
		} else if tmpNode.Next.Data >= node.Data {
			node.Next = tmpNode.Next
			tmpNode.Next = node
			break
		}
		tmpNode = tmpNode.Next
	}
	return root
}

func DeleteLink(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	if tmpNode != nil && tmpNode.Data == data {
		if tmpNode.Next == nil {
			return nil
		}
		rightNode := tmpNode.Next
		tmpNode.Next = nil
		return rightNode
	}
	for {
		if tmpNode.Next == nil {
			break
		}
		rightNode := tmpNode.Next
		if rightNode.Data == data {
			tmpNode.Next = rightNode.Next
			rightNode.Next = nil
			return root
		}
		tmpNode = tmpNode.Next
	}
	return root
}

func rangeLink(root *LinkNode) {
	for root != nil {
		fmt.Println(root.Data)
		if root.Next != nil {
			root = root.Next
		} else {
			break
		}
	}
}

func main() {
	n1 := &LinkNode{
		Data: 1,
		Next: nil,
	}

	n2 := &LinkNode{
		Data: 2,
		Next: nil,
	}

	n3 := &LinkNode{
		Data: 3,
		Next: nil,
	}

	n4 := &LinkNode{
		Data: 4,
		Next: nil,
	}

	n6 := &LinkNode{
		Data: 6,
		Next: nil,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n6

	n5 := &LinkNode{
		Data: 5,
		Next: nil,
	}
	tmpNode := n1
	tmpNode = InsertLink(tmpNode, n5)
	tmpNode = InsertLink(tmpNode, &LinkNode{
		Data: 6,
		Next: nil,
	})
	tmpNode = InsertLink(tmpNode, &LinkNode{
		Data: 7,
		Next: nil,
	})
	tmpNode = InsertLink(tmpNode, &LinkNode{
		Data: 0,
		Next: nil,
	})
	tmpNode = InsertLink(tmpNode, &LinkNode{
		Data: 0,
		Next: nil,
	})
	rangeLink(tmpNode)
	nl := DeleteLink(tmpNode, 0)
	nl = DeleteLink(nl, 0)
	rangeLink(nl)
}
