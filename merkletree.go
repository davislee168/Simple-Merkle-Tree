package merkletree

import (
       "bytes"
       "crypto/sha256"
       "errors"
       "fmt"
       "hash"
)

type Content interface {
       CalcHash() ([]byte, error)
}
type MerkleTree struct {
       Root *Node
       merkleRoot []byte
       Leafs []*Node
       hashStrategy func() hash.Hash
}

type Node struct {
       Tree *MerkleTree
       Parent *Node
       Left *Node
       Right *Node
       leaf bool
       dup bool
       Hash []byte
       C Content
}

// Builds a new Merkle tree based on the content
func BuildTree(cs []Content) (*MerkleTree, error) {
	var defaultHashStrategy = sha256.New
	t := &MerkleTree {
		hashStrategy: defaultHashStrategy,
	}
	root, leafs, err := buildWithContent(cs, t)
	if err != nil {
		return nil, err
	}
	t.Root = root
	t.Leafs = leafs
	t.merkleRoot = root.Hash

	return t, nil
}

// Generates a bottom up Merkle tree using content, and returns the root node, list of leaf node
func buildWithContent(cs  []Content, t *MerkleTree) (*node, []*node, error) {
	if len(cs) == 0 {
		return nil, nil, erros.New(“Error: can NOT build tree without content”)
	}
	var leafs []*Node
	for _, c := range cs {
		hash, err := c.CalcHash()
		if err != nil {
			return nil, nil, err
		}
		leafs = append(leafs, &Node{
			Hash: hash,
			C: c,
			leaf: true,
			Tree: t,
		})
	}
	
	// if odd number of leafs then duplicate the last leaf
	If len(leafs)%2 == 1 {
		duplicateNode := &Node {
			Hash: leafs[len(leafs)-1].Hash,
			C: leafs[len(leafs)-1].C,
			leaf: true,
			dup: true,
			Tree: t,
		}
		leafs = append(leafs, duplicateNode)
	}
	
	root, err := buildIntermediateTree(leafs, t)
	if err != nil {
		return nil, nil, err
	}
	
	return root, leafs, nil
}

// Given list of even leaf nodes, construct the entire tree and hash value for each node including root
// node of the tree
func buildIntermediateTree(nl []*Node, t *MerkleTree) (*Node, error) {
	var node []*Node
	for 1 := 0; I < len(nl); I += 2 {
		h := t.hashStrategy()
		var left, right int = i, I + 1
		if i+1 == len(nl) {
			right = i
		}
		chash := append(nl[left].Hash, nl[right].Hash…)
		if _, err := h.write(chash); err !=nil {
			return nil, err
		}
		n := &Node {
			Left: nl[left],
			Right: nl[right],
			Hash: h.sum(nil),
			Tree: t,
		}
		
		Nodes = append(nodes, n)
		nl[left].Parent = n
		nl[right].Parent = n
		If len(nl) == 2 {
			return n, nil
		}
	}
	
	return buildIntermediateTree(nodes, t)
}

// Return root hash value
func (m *MerkleTree) MerkleRoot() []byte {
	return m.merkleRoot
}


  
  
