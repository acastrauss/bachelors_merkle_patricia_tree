package models

func (key NodeKey) getPrefix(nodeType NodeType) Prefix {
	lengthModule := key.getLength() % 2
	return Prefix(lengthModule + int(nodeType)*2)
}

func (key NodeKey) getLength() int {
	return len(key.Key)
}

func (key NodeKey) equals(refKey NodeKey) bool {
	return key.Key == refKey.Key
}
