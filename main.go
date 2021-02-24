package main

import "fmt"

type Node struct {
	No           string
	ParentNode   *Node
	SubNodeSlice []*Node
}

// var AllNodeMap map[string]*Node
var RootNodeMap map[string]*Node
var NonRootNodeMap map[string]*Node
var WhiteCollection map[string]*Node
var GreyCollection map[string]*Node
var BlackCollection map[string]*Node

// simulate Go GC algorithm in v1.5

func main() {
	Step0InitNodeTree()
	Step1ContributeCollection()
	Step2MarkAllNodesWhite()
	Step3ScanRootNodeSliceAndMarkGrey()
	Step4ScanGreyCollectionAndMarkBlack()
	Step5SweepWhiteCollection()
}

func Step0InitNodeTree() {
	ANode := &Node{No: "A"}
	BNode := &Node{No: "B"}
	CNode := &Node{No: "C"}
	DNode := &Node{No: "D"}
	ENode := &Node{No: "E"}
	FNode := &Node{No: "F"}
	GNode := &Node{No: "G"}
	HNode := &Node{No: "H"}

	RootNode1 := &Node{No: "Root1"}
	RootNode2 := &Node{No: "Root2"}
	RootNode3 := &Node{No: "Root3"}

	RootNodeMap = map[string]*Node{
		RootNode1.No: RootNode1, RootNode2.No: RootNode2, RootNode3.No: RootNode3,
	}

	NonRootNodeMap = map[string]*Node{
		ANode.No: ANode, BNode.No: BNode, CNode.No: CNode, DNode.No: DNode,
		ENode.No: ENode, FNode.No: FNode, GNode.No: GNode, HNode.No: HNode,
	}

	RootNode1.SubNodeSlice = append(RootNode1.SubNodeSlice, ANode)
	ANode.ParentNode = RootNode1
	ANode.SubNodeSlice = append(ANode.SubNodeSlice, BNode)
	BNode.ParentNode = ANode
	ANode.SubNodeSlice = append(ANode.SubNodeSlice, CNode)
	CNode.ParentNode = ANode
	ANode.SubNodeSlice = append(ANode.SubNodeSlice, DNode)
	DNode.ParentNode = ANode
	// circular reference
	DNode.SubNodeSlice = append(DNode.SubNodeSlice, ANode)

	RootNode3.SubNodeSlice = append(RootNode3.SubNodeSlice, FNode)
	FNode.ParentNode = RootNode3

	GNode.SubNodeSlice = append(GNode.SubNodeSlice, HNode)
	HNode.ParentNode = GNode

	// Root1 Root2 Root3 ...
	//   |           |
	//   A           F   E   G
	//  /|\                  |
	// B C D                 H
}

func Step1ContributeCollection() {
	WhiteCollection, GreyCollection, BlackCollection = make(map[string]*Node), make(map[string]*Node), make(map[string]*Node)
}

// Step2MarkAllNodesWhite
// A...H mark White
func Step2MarkAllNodesWhite() {
	for no, node := range NonRootNodeMap {
		WhiteCollection[no] = node
	}
}

// Step3ScanRootNodeSliceAndMarkGrey
// A F mark Grey
// A F remove White
func Step3ScanRootNodeSliceAndMarkGrey() {
	for _, node := range RootNodeMap {
		for _, subNode := range node.SubNodeSlice {
			GreyCollection[subNode.No] = subNode
			delete(WhiteCollection, subNode.No)
		}
	}
}

// Step4ScanGreyCollection
// B C D mark Grey, B C D remove White, A F mark Black Collection
// B C D mark Black Collection
func Step4ScanGreyCollectionAndMarkBlack() {
	for len(GreyCollection) > 0 {
		newGreyCollection := make(map[string]*Node)
		for no, node := range GreyCollection {
			BlackCollection[no] = node
			for _, subNode := range node.SubNodeSlice {
				if _, isBlack := BlackCollection[subNode.No]; !isBlack {
					newGreyCollection[subNode.No] = subNode
					delete(WhiteCollection, subNode.No)
				}
			}
		}
		GreyCollection = newGreyCollection
	}
}

// Step5SweepWhiteCollection
// Sweep Node: E G H
func Step5SweepWhiteCollection() {
	for _, node := range WhiteCollection {
		fmt.Printf("sweep node %v\n", node.No)
	}
}
