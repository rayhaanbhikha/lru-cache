package linkedlist

import "fmt"

type LinkList struct {
	size             int
	startingNullNode *LinkedListNode
	endNullNode      *LinkedListNode
}

func (li *LinkList) String() string {
	return li.startingNullNode.String()
}

func New() *LinkList {
	startingNullNode, endNullNode := &LinkedListNode{}, &LinkedListNode{}
	endNullNode.prev = startingNullNode
	startingNullNode.next = endNullNode
	return &LinkList{startingNullNode: startingNullNode, endNullNode: endNullNode}
}

func (li *LinkList) AddNode(newNode *LinkedListNode) *LinkedListNode {
	li.startingNullNode.addNode(newNode)
	li.size++
	return newNode
}

func (li *LinkList) Add(val int) *LinkedListNode {
	newNode := &LinkedListNode{value: val}
	li.startingNullNode.addNode(newNode)
	li.size++
	return newNode
}

func (li *LinkList) Pop() *LinkedListNode {
	if li.size == 0 {
		return nil
	}
	return li.endNullNode.prev.Delete()
}

type LinkedListNode struct {
	value int
	prev  *LinkedListNode
	next  *LinkedListNode
}

func (n *LinkedListNode) addNode(newNode *LinkedListNode) {
	// We will only be adding from the front.
	nextNode := n.next

	n.next = newNode
	newNode.next = nextNode
	newNode.prev = n

	nextNode.prev = newNode
}

func (n *LinkedListNode) Delete() *LinkedListNode {
	if n.next == nil {
		// This shouldn't have happened.
		return nil
	}
	nextNode := n.next
	n.prev.next = nextNode
	nextNode.prev = n.prev
	return n
}

func (n *LinkedListNode) GetValue() int {
	return n.value
}

func (n *LinkedListNode) String() string {
	head := n
	s := fmt.Sprintf("%v -> ", head.value)
	for head.next != nil {
		head = head.next
		s += fmt.Sprintf("%v -> ", head.value)
	}
	return s
}
