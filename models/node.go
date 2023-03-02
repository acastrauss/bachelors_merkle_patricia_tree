package models

type Any interface{}

type Prefix int

const (
	ExtensionEven Prefix = 0
	ExtensionOdd  Prefix = 1
	LeafEven      Prefix = 2
	LeafOdd       Prefix = 3
)

type NodeType int

const (
	Extension NodeType = 0
	Leaf      NodeType = 1
)

type NodeKey struct {
	Key string
}

type NodeValue struct {
	Value int
}

type NodeWithKey interface {
	GetKey() NodeKey
	GetPrefix() Prefix
}

type ExtensionNode struct {
	NodePrefix Prefix
	SharedKey  NodeKey
	ChildPtr   Any
	ParentPtr  Any
}

func (en ExtensionNode) GetKey() NodeKey {
	return en.SharedKey
}

func (en ExtensionNode) GetPrefix() Prefix {
	return en.NodePrefix
}

type BranchNode struct {
	ChildPtrs []Any
	Value     NodeValue
	ParentPtr Any
}

const (
	NUM_OF_BRANCH_CHILDREN = 16
)

type LeafNode struct {
	NodePrefix Prefix
	KeyEnd     NodeKey
	Value      NodeValue
	ParentPtr  *BranchNode
}

func (ln LeafNode) GetKey() NodeKey {
	return ln.KeyEnd
}

func (ln LeafNode) GetPrefix() Prefix {
	return ln.NodePrefix
}

type MPT struct {
	Root NodeWithKey
}
