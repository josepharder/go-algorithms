package linkedList

import "errors"

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
