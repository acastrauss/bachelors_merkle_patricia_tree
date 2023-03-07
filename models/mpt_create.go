package models

func CreateMPT(key NodeKey, value NodeValue) MPT {
	return MPT{
		Root: LeafNode{
			KeyEnd: key,
			Value:  value,
			Prefix: key.getPrefix(Leaf),
			Parent: nil,
		},
	}
}

func (mpt *MPT) InsertKVPair(key NodeKey, value NodeValue) {
	lsm := GetLastSimilarNode(mpt.Root, key)
	if lsm.Node == nil {
		// true only when there is 1 node in trie
		lsm.Node = &mpt.Root
	}

	if lastLeaf, ok := (*lsm.Node).(*LeafNode); ok {
		// make extension from leaf, and branch as a child
		sharedKey := lastLeaf.GetKey().getSharedKeyWithGivenKey(key)

		newExtension := ExtensionNode{
			SharedKey:   sharedKey,
			Prefix:      sharedKey.getPrefix(Extension),
			BranchChild: nil,
			Parent:      lastLeaf.GetParent(),
		}

		newBranch := BranchNode{
			Children: make(map[rune]*TrieNode),
			// Parent:   &TrieNode(&newExtension),
			Value: NodeValue{Value: 0},
		}

		*newBranch.Parent = newExtension
		*newExtension.BranchChild = newBranch

		lsnKeyLen := lastLeaf.GetKey().getLength()
		leaf1Key := NodeKey{Key: lastLeaf.GetKey().Key[lsnKeyLen-lsm.KeyDifference.getLength():]}
		leaf2Key := lsm.KeyDifference

		*(newBranch.Children[rune(leaf1Key.Key[0])]) = LeafNode{
			KeyEnd: NodeKey{Key: leaf1Key.Key[1:]},
			Prefix: NodeKey{Key: leaf1Key.Key[1:]}.getPrefix(Leaf),
			Value:  lastLeaf.Value,
			// Parent: newBranch,
		}
		tempChildPtr := newBranch.Children[rune(leaf1Key.Key[0])]
		*((*tempChildPtr).(LeafNode).Parent) = newBranch

		*(newBranch.Children[rune(leaf2Key.Key[0])]) = LeafNode{
			KeyEnd: NodeKey{Key: leaf2Key.Key[1:]},
			Prefix: NodeKey{Key: leaf2Key.Key[1:]}.getPrefix(Leaf),
			Value:  value,
			// Parent: newBranch,
		}

		temp2ChildPtr := newBranch.Children[rune(leaf2Key.Key[0])]
		*((*temp2ChildPtr).(LeafNode).Parent) = newBranch

	} else if lastBranch, ok := (*lsm.Node).(*BranchNode); ok {
		if lastBranch.IsKeyInBranch(lsm.KeyDifference) {
			//nodeAtLastBranch := lastBranch.GetNodeAt(rune( /*last similar rune from key for indexing branch*/ ))
			panic("Then there is node that is after branch")
		} else {
			newLeaf := LeafNode{
				KeyEnd: lsm.KeyDifference,
				Prefix: lsm.KeyDifference.getPrefix(Leaf),
				Value:  value,
				// Parent: lastSimilarNode.Node, // parent will be current branch
			}

			*newLeaf.Parent = lastBranch
			*(lastBranch.Children[rune(lsm.KeyDifference.Key[0])]) = newLeaf
		}
	} else {
		panic("last node can not be extension")
	}
}

type LastSimilarNode struct {
	Node            *TrieNode
	KeyDifference   NodeKey
	LastSimilarRune rune
}

func GetLastSimilarNode(nodeToCompareTo TrieNode, key NodeKey) LastSimilarNode {
	tearedApartKey := nodeToCompareTo.TearApartGivenKeyWithMine(key)
	if nodeToCompareTo.GetType() == Branch {
		// compare only first element of 'key'
		branch := nodeToCompareTo.(BranchNode)
		if branch.IsKeyInBranch(key) {
			return GetLastSimilarNode(*branch.GetNodeAt(rune(key.Key[0])), tearedApartKey)
		} else {
			return LastSimilarNode{
				Node:            &nodeToCompareTo,
				KeyDifference:   tearedApartKey,
				LastSimilarRune: nodeToCompareTo.GetLastSimilarRuneWithMyKey(tearedApartKey),
			}
		}
	} else {
		if nodeToCompareTo.HasChildren() {
			return GetLastSimilarNode(*nodeToCompareTo.(ExtensionNode).BranchChild, tearedApartKey)
		} else {

			similarNode := nodeToCompareTo

			if nodeToCompareTo.GetKey().indexOfSimilarityWithGivenKey(key) == NO_SIMILARITY {
				similarNode = *nodeToCompareTo.GetParent()
			}
			return LastSimilarNode{
				Node:            &similarNode,
				KeyDifference:   tearedApartKey,
				LastSimilarRune: nodeToCompareTo.GetLastSimilarRuneWithMyKey(tearedApartKey),
			}
		}
	}
}
