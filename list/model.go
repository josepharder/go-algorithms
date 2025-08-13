package list

type LinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}
