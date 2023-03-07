package models

type LeafNode struct {
	Prefix NodePrefix
	KeyEnd NodeKey
	Value  NodeValue
	Parent *TrieNode
}

func (ln LeafNode) GetKey() NodeKey {
	return ln.KeyEnd
}

func (ln *LeafNode) SetKey(key NodeKey) {
	ln.KeyEnd = key
}

func (ln LeafNode) TearApartGivenKeyWithMine(key NodeKey) NodeKey {
	return ln.KeyEnd.tearApartGivenKeyWithMe(key)
}

func (ln LeafNode) GetLastSimilarRuneWithMyKey(key NodeKey) rune {
	return ln.KeyEnd.getLastSimilarRuneWithGivenKey(key)
}

func (ln LeafNode) GetValue() NodeValue {
	return ln.Value
}

func (ln *LeafNode) SetValue(value NodeValue) {
	ln.Value = value
}

func (ln LeafNode) GetPrefix() NodePrefix {
	return ln.Prefix
}

func (ln *LeafNode) SetPrefix(prefix NodePrefix) {
	if prefix != LeafEven && prefix != LeafOdd {
		panic("wrong prefix for a leaf node")
	}

	ln.Prefix = prefix
}

func (ln LeafNode) GetType() NodeType {
	return Leaf
}

func (ln LeafNode) GetChildren() map[rune]TrieNode {
	panic("leaf node doesn't have any children")
}

func (ln LeafNode) SetChildren(_ map[rune]TrieNode) {
	panic("leaf node doesn't have any children")
}

func (ln LeafNode) GetParent() *TrieNode {
	return ln.Parent
}

func (ln LeafNode) HasChildren() bool {
	return false
}

func (ln *LeafNode) SetParent(parent *TrieNode) {
	ln.Parent = parent
}
