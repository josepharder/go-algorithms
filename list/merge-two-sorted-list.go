package list

var list1 LinkedList
var list2 LinkedList

func p() {
	list1.Append(1)
	list1.Append(2)
	list1.Append(4)

	list2.Append(0)
	list2.Append(3)
	list2.Append(4)

	mergeTwoLists(list1.Head, list2.Head)
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	isList1Empty := list1 == nil
	isList2Empty := list2 == nil

	if isList1Empty && isList2Empty {
		return nil
	}
	if isList1Empty {
		return list2
	}

	if isList2Empty {
		return list1
	}

	newList := &Node{}
	current := newList
	for list1 != nil && list2 != nil {

		if list1.Value <= list2.Value {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return newList.Next
}
