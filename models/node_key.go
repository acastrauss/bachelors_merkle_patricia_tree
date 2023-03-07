package models

type NodeKey struct {
	Key string
}

func (key NodeKey) getPrefix(nodeType NodeType) NodePrefix {
	lengthModule := key.getLength() % 2
	return NodePrefix(lengthModule + int(nodeType)*2)
}

func (key NodeKey) getLength() int {
	return len(key.Key)
}

func (key NodeKey) equals(refKey NodeKey) bool {
	return key.Key == refKey.Key
}

const (
	NO_SIMILARITY = -1
)

func (thisKey NodeKey) indexOfSimilarityWithGivenKey(key NodeKey) int {
	indexOfSimilarity := -1

	for i := 0; i < key.getLength() && i < thisKey.getLength(); i++ {
		if thisKey.Key[i] == key.Key[i] {
			indexOfSimilarity++
		} else {
			break
		}
	}

	return indexOfSimilarity
}

func (thisKey NodeKey) tearApartGivenKeyWithMe(keyToTearApart NodeKey) NodeKey {
	if thisKey.getLength() > keyToTearApart.getLength() {
		panic("can not tear apart smaller key with a larger one")
	}

	indexOfSimilarity := 0

	for i := 0; i < keyToTearApart.getLength(); i++ {
		if thisKey.Key[i] == keyToTearApart.Key[i] {
			indexOfSimilarity++
		} else {
			break
		}
	}

	return NodeKey{
		Key: keyToTearApart.Key[indexOfSimilarity:],
	}
}

func (thisKey NodeKey) getLastSimilarRuneWithGivenKey(key NodeKey) rune {
	lastSimilarIndx := -1

	for i := 0; i < thisKey.getLength() && i < key.getLength(); i++ {
		if thisKey.Key[i] == key.Key[i] {
			lastSimilarIndx = i
		} else {
			break
		}
	}

	if lastSimilarIndx == -1 {
		panic("no similarity")
	}

	return rune(key.Key[lastSimilarIndx])
}

func (thisKey NodeKey) getSharedKeyWithGivenKey(key NodeKey) NodeKey {
	sharedKey := NodeKey{Key: ""}
	for i := 0; i < key.getLength(); i++ {
		if thisKey.Key[i] == key.Key[i] {
			sharedKey.Key = string(sharedKey.Key + string(thisKey.Key[i]))
		} else {
			break
		}
	}

	if sharedKey.getLength() == 0 {
		panic("no shared key")
	}

	return sharedKey
}
