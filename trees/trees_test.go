package trees

import (
	"fmt"
	"testing"
)

func TestTrees(t *testing.T) {
	company := NewTreeNode("company")
	cto := NewTreeNode("cto")

	company.AddChild(cto)

	eng := NewTreeNode("eng")
	mrk := NewTreeNode("marketing")
	ops := NewTreeNode("operations")

	cto.AddChild(eng)
	cto.AddChild(mrk)
	cto.AddChild(ops)

	senior_eng := NewTreeNode("senior eng")
	senior_qa := NewTreeNode("senior qa")

	dev := NewTreeNode("dev")
	qa := NewTreeNode("qa")

	eng.AddChild(senior_eng)
	eng.AddChild(senior_qa)

	senior_eng.AddChild(dev)
	senior_qa.AddChild(qa)

	PrintTree(cto, "", false)

	// finding a node
	fmt.Println("Is qa found?", FindNode(cto, "eng"))

	//tree height
	fmt.Println("Tree height", GetHeight(company))

	//leaves of a tree
	fmt.Println("Leaves of a tree", GetLeaves(company))

	// Path of a tree

	fmt.Println("Path to senior QA", GetPath(company, "senior qa"))

	// Depth first traversal
	DFT(company, func(s string) {
		fmt.Printf("%s, ", s)
	})

	fmt.Println()

	// Bread first traversal
	BFT(company, func(s string) {
		fmt.Printf("%s, ", s)
	})
}
