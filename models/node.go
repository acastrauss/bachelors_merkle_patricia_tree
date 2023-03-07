package models

type Any interface{}

type NodePrefix int

const (
	ExtensionEven NodePrefix = 0
	ExtensionOdd  NodePrefix = 1
	LeafEven      NodePrefix = 2
	LeafOdd       NodePrefix = 3
	NoPrefix      NodePrefix = -1
)

type NodeType int

const (
	Extension NodeType = 0
	Leaf      NodeType = 1
	Branch    NodeType = 2
)

type NodeValue struct {
	Value int
}

type TrieNode interface {
	GetKey() NodeKey
	TearApartGivenKeyWithMine(NodeKey) NodeKey
	GetLastSimilarRuneWithMyKey(NodeKey) rune

	GetValue() NodeValue

	GetPrefix() NodePrefix

	GetType() NodeType

	HasChildren() bool

	GetParent() *TrieNode
}

type MPT struct {
	Root TrieNode
}
