package models

import (
	"reflect"
	"strconv"
)

func CreateEmptyMPT() MPT {
	return MPT{
		Root: nil,
	}
}

const (
	NO_SIMILARITY = -1
)

func (mpt *MPT) InsertKVPair(key NodeKey, value NodeValue) {
	if mpt.Root == nil {
		mpt.Root = LeafNode{
			NodePrefix: key.getPrefix(Leaf),
			KeyEnd:     key,
			Value:      value,
			ParentPtr:  nil,
		}
	} else {
		ios := GetIndexOfSimilarity(mpt.Root.GetKey(), key)

		if ios == NO_SIMILARITY || key.getLength() == ios {
			panic("Key is not for this tree")
		}

		if reflect.TypeOf(mpt.Root) == reflect.TypeOf(ExtensionNode{}) {
			tempBranch := mpt.Root.(ExtensionNode).ChildPtr.(BranchNode)
			branchIndx, _ := strconv.ParseInt(string(key.Key[ios+1]), 16, 32)
			tempBranch.AddToBranch(key, value, int(branchIndx), ios+1, mpt, ios)
			// mpt.Root.(*ExtensionNode).ChildPtr = tempBranch
		} else {
			mpt.SplitRoot(key, value, ios)
		}
	}
}

func (bn BranchNode) AddToBranch(key NodeKey, value NodeValue, currentBranchIndx int, currentKeyIndx int, mpt *MPT, ios int) {
	if bn.ChildPtrs[currentBranchIndx] != nil {

		if reflect.TypeOf(bn.ChildPtrs[currentBranchIndx]) == reflect.TypeOf(BranchNode{}) {
			currentKeyIndx += 1
			branchIndx, _ := strconv.ParseInt(string(key.Key[ios+currentKeyIndx]), 16, 32)
			bn.ChildPtrs[currentBranchIndx].(BranchNode).AddToBranch(key, value, int(branchIndx), currentKeyIndx, mpt, ios)
		} else {
			tempMpt := MPT{
				Root: bn.ChildPtrs[currentBranchIndx].(NodeWithKey),
			}
			newios := GetIndexOfSimilarity(NodeKey{Key: key.Key[currentKeyIndx+1:]}, bn.ChildPtrs[currentBranchIndx].(NodeWithKey).GetKey())
			tempMpt.SplitRoot(NodeKey{Key: key.Key[currentKeyIndx+1:]}, value, newios)

			bn.ChildPtrs[currentBranchIndx] = tempMpt.Root
		}
	} else if reflect.TypeOf((mpt.Root.(ExtensionNode).ChildPtr)) == reflect.TypeOf(BranchNode{}) && reflect.TypeOf(bn.ParentPtr) == reflect.TypeOf(ExtensionNode{}) {
		// split current root, but one of new children will be branch
		indx, _ := strconv.ParseInt(string(mpt.Root.GetKey().Key[ios+1]), 16, 32)

		mpt.SplitRootWithBranchNode(key, value, ios, bn, int(indx))
	} else {
		leafKey := NodeKey{
			Key: key.Key[currentKeyIndx:],
		}
		newLeaf := LeafNode{
			ParentPtr:  &bn,
			Value:      value,
			NodePrefix: leafKey.getPrefix(Leaf),
			KeyEnd:     leafKey,
		}
		bn.ChildPtrs[currentBranchIndx] = newLeaf
	}
}

func (mpt *MPT) SplitRoot(key NodeKey, value NodeValue, ios int) {
	rootSharedKey := NodeKey{
		Key: key.Key[0 : ios+1],
	}
	newRoot := ExtensionNode{
		ParentPtr:  nil,
		NodePrefix: rootSharedKey.getPrefix(Extension),
		SharedKey:  rootSharedKey,
		ChildPtr:   nil,
	}

	child1Key := NodeKey{
		Key: key.Key[ios+2:],
	}
	child1 := LeafNode{
		NodePrefix: child1Key.getPrefix(Leaf),
		KeyEnd:     child1Key,
		Value:      value,
		ParentPtr:  nil,
	}

	child2Key := NodeKey{
		Key: mpt.Root.GetKey().Key[ios+2:],
	}
	child2 := LeafNode{
		NodePrefix: child2Key.getPrefix(Leaf),
		KeyEnd:     child2Key,
		Value:      NodeValue{Value: -1111111},
		ParentPtr:  nil,
	}

	newBranch := BranchNode{
		Value:     NodeValue{},
		ParentPtr: newRoot,
		ChildPtrs: make([]Any, NUM_OF_BRANCH_CHILDREN),
	}

	child1.ParentPtr = &newBranch
	child2.ParentPtr = &newBranch

	indx1, _ := strconv.ParseInt(string(key.Key[ios+1]), 16, 32)
	newBranch.ChildPtrs[indx1] = child1
	indx2, _ := strconv.ParseInt(string(mpt.Root.GetKey().Key[ios+1]), 16, 32)
	newBranch.ChildPtrs[indx2] = child2

	newBranch.ParentPtr = newRoot
	newRoot.ChildPtr = newBranch

	mpt.Root = newRoot
}

func (mpt *MPT) SplitRootWithBranchNode(key NodeKey, value NodeValue, ios int, bn BranchNode, indxForNewBN int) {
	rootSharedKey := NodeKey{
		Key: key.Key[0 : ios+1],
	}
	newRoot := ExtensionNode{
		ParentPtr:  nil,
		NodePrefix: rootSharedKey.getPrefix(Extension),
		SharedKey:  rootSharedKey,
		ChildPtr:   nil,
	}

	child1Key := NodeKey{
		Key: key.Key[ios+2:],
	}
	child1 := LeafNode{
		NodePrefix: child1Key.getPrefix(Leaf),
		KeyEnd:     child1Key,
		Value:      value,
		ParentPtr:  nil,
	}

	newBranch := BranchNode{
		Value:     NodeValue{},
		ParentPtr: newRoot,
		ChildPtrs: make([]Any, NUM_OF_BRANCH_CHILDREN),
	}

	child1.ParentPtr = &newBranch
	bn.ParentPtr = &newBranch

	indx1, _ := strconv.ParseInt(string(key.Key[ios+1]), 16, 32)
	newBranch.ChildPtrs[indx1] = child1
	newBranch.ChildPtrs[indxForNewBN] = bn

	newBranch.ParentPtr = newRoot
	newRoot.ChildPtr = newBranch

	mpt.Root = newRoot
}

func GetIndexOfSimilarity(key1 NodeKey, key2 NodeKey) int {
	if key1.equals(key2) {
		panic("same keys")
	}

	indexOfSimilarity := NO_SIMILARITY

	for i := 0; i < key1.getLength() && i < key2.getLength(); i++ {
		if key1.Key[i] == key2.Key[i] {
			indexOfSimilarity++
		} else {
			break
		}
	}

	return indexOfSimilarity
}
