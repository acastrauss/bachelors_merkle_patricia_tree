package models

const (
	NUM_OF_BRANCH_CHILDREN = 16
)

type BranchNode struct {
	Children map[rune]*TrieNode
	Parent   *TrieNode
	Value    NodeValue
}

func (bn BranchNode) GetNodeAt(indx rune) *TrieNode {
	return bn.Children[indx]
}

func (bn BranchNode) GetKey() NodeKey {
	return NodeKey{
		Key: "0123456789abcdef",
	}
}

func (bn BranchNode) IsKeyInBranch(key NodeKey) bool {
	_, inBranch := bn.Children[rune(key.Key[0])]
	return inBranch
}

func (bn BranchNode) SetKey(_ NodeKey) {
	panic("branch node doesn't have settable key")
}

func (bn BranchNode) TearApartGivenKeyWithMine(key NodeKey) NodeKey {
	if bn.IsKeyInBranch(key) {
		return NodeKey{
			Key: key.Key[1:],
		}
	} else {
		return key
	}
}

func (bn BranchNode) GetLastSimilarRuneWithMyKey(key NodeKey) rune {
	if bn.IsKeyInBranch(NodeKey{Key: string(key.Key[0])}) {
		return rune(key.Key[0])
	} else {
		panic("key doesn't belong in a branch")
	}
}

func (bn BranchNode) GetValue() NodeValue {
	return bn.Value
}

func (bn *BranchNode) SetValue(value NodeValue) {
	bn.Value = value
}

func (bn BranchNode) GetPrefix() NodePrefix {
	panic("branch node doesn't have a prefix")
}

func (bn BranchNode) SetPrefix(_ NodePrefix) {
	panic("branch node doesn't have a prefix")
}

func (bn BranchNode) GetType() NodeType {
	return Branch
}

func (bn BranchNode) HasChildren() bool {
	for _, v := range bn.Children {
		if v != nil {
			return true
		}
	}

	return false
}

func (bn BranchNode) GetParent() *TrieNode {
	return bn.Parent
}

func (bn *BranchNode) SetParent(parent *TrieNode) {
	bn.Parent = parent
}

func (bn *BranchNode) InsertAt(indx rune, node *TrieNode) {
	bn.Children[indx] = node
}
