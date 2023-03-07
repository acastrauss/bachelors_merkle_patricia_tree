package models

const (
	EXTENSION_NODE_CHILDREN_INDX = '0'
)

type ExtensionNode struct {
	Prefix      NodePrefix
	SharedKey   NodeKey
	BranchChild *TrieNode
	Parent      *TrieNode
}

func (en ExtensionNode) GetKey() NodeKey {
	return en.SharedKey
}

func (en *ExtensionNode) SetKey(k NodeKey) {
	en.SharedKey = k
}

func (en ExtensionNode) TearApartGivenKeyWithMine(key NodeKey) NodeKey {
	return en.SharedKey.tearApartGivenKeyWithMe(key)
}

func (en ExtensionNode) GetLastSimilarRuneWithMyKey(key NodeKey) rune {
	return en.SharedKey.getLastSimilarRuneWithGivenKey(key)
}

func (en ExtensionNode) GetValue() NodeValue {
	panic("extension node doesn't have a value")
}

func (en *ExtensionNode) SetValue(_ NodeValue) {
	panic("extension node doesn't have a value")
}

func (en ExtensionNode) GetPrefix() NodePrefix {
	return en.Prefix
}

func (en *ExtensionNode) SetPrefix(prefix NodePrefix) {
	if prefix != ExtensionEven && prefix != ExtensionOdd {
		panic("wrong prefix for an extension node")
	}
	en.Prefix = prefix
}

func (en ExtensionNode) GetType() NodeType {
	return Extension
}

func (en ExtensionNode) HasChildren() bool {
	return en.BranchChild != nil
}

func (en ExtensionNode) GetParent() *TrieNode {
	return en.Parent
}

func (en *ExtensionNode) SetParent(parent *TrieNode) {
	en.Parent = parent
}
