package trees

import "fmt"

type TreeNode struct {
	Value    string
	Children []*TreeNode
}

func NewTreeNode(value string) *TreeNode {
	return &TreeNode{
		Value:    value,
		Children: make([]*TreeNode, 0),
	}
}

func (t *TreeNode) AddChild(child *TreeNode) {
	t.Children = append(t.Children, child)
}

func (t *TreeNode) RemoveChild(target string) bool {
	for i, child := range t.Children {
		if child.Value == target {
			t.Children = append(t.Children[:i], t.Children[i+1:]...)
			return true
		}
	}

	return false
}

func (t *TreeNode) IsLeaf() bool {
	return len(t.Children) == 0
}

func (t *TreeNode) GetChildCount() int {
	count := 0

	for _, child := range t.Children {
		count += child.GetChildCount()
	}

	return count
}

func (t *TreeNode) HasChild(value string) bool {
	for _, child := range t.Children {
		if child.Value == value {
			return true
		}
	}

	return false
}

func FindNode(root *TreeNode, target string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Value == target {
		return root
	}

	for _, child := range root.Children {
		if found := FindNode(child, target); found != nil {
			return found
		}
	}

	return nil
}

func GetHeight(node *TreeNode) int {
	if len(node.Children) == 0 || node == nil {
		return 0
	}

	maxHeight := 0

	for _, child := range node.Children {
		height := GetHeight(child)
		if height > maxHeight {
			maxHeight = height
		}

	}

	return maxHeight + 1
}

func GetLeaves(node *TreeNode) []string {
	if node == nil {
		return []string{}
	}

	var leaves []string
	if node.IsLeaf() {
		return []string{node.Value}
	}

	for _, child := range node.Children {
		leaves = append(leaves, GetLeaves(child)...)
	}

	return leaves
}

func GetPath(root *TreeNode, target string) []string {
	if root == nil {
		return nil
	}

	if root.Value == target {
		return []string{root.Value}
	}

	for _, child := range root.Children {
		if path := GetPath(child, target); path != nil {
			return append([]string{root.Value}, path...)
		}
	}

	return nil
}

func GetNodesAtLevel(node *TreeNode, targetLevel int) []string {
	if node == nil {
		return nil
	}

	if targetLevel == 0 {
		return []string{node.Value}
	}

	var nodes []string
	for _, child := range node.Children {
		nodes = append(nodes, GetNodesAtLevel(child, targetLevel-1)...)
	}

	return nodes
}

func DFT(node *TreeNode, visit func(string)) {
	if node == nil {
		return
	}

	visit(node.Value)

	for _, child := range node.Children {
		DFT(child, visit)
	}
}

func BFT(node *TreeNode, visit func(string)) {
	if node == nil {
		return
	}

	queue := []*TreeNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visit(current.Value)

		queue = append(queue, current.Children...)
	}

}

func PrintTree(node *TreeNode, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	fmt.Printf("%s%s%s\n", prefix, connector, node.Value)

	nextPrefix := prefix
	if isLast {
		nextPrefix += "    "
	} else {
		nextPrefix += "│   "
	}

	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		PrintTree(child, nextPrefix, isLastChild)
	}
}
