package linkedList

import (
	"errors"
)

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	l.Size++
}

func (l *LinkedList) RemoveNodeByValue(value int) error {
	if l.Size == 0 {
		return errors.New("list is empty")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Size--
		return nil
	}

	// find the prev node to node to remove
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.Size--

	return nil
}

func (l *LinkedList) RemoveNodeByPosition(position int) error {
	if l.Size == 0 {
		return errors.New("list is empty")
	}

	if position == 0 {
		l.Head = l.Head.Next
		l.Size--
		return nil
	}

	current := l.Head

	for i := 0; i < position-1; i++ {
		if current.Next != nil {
			current = current.Next
		}
	}

	current.Next = current.Next.Next
	l.Size--

	return nil
}

func (l *LinkedList) FindMiddleNode() (*Node, error) {

	if l.Head == nil {
		return nil, errors.New("list is empty")
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow, nil
}
