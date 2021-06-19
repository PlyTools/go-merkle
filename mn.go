package go_merkle

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Data interface {
	CalculateHash() ([]byte, error)
	Equals(Data) (bool, error)
}

// A MerkleNode represents a merkle tree node
//
// For leaf nodes, the Children should be nil while the Data is not
// For non-leaf nodes, the Data should be nil while the Children are not
type MerkleNode struct {
	Children []*MerkleNode // allow Multiway Trees
	Hash     []byte  // allow different hash functions
	Data 	 Data	// allow different data types
}

func NewMerkleNode(children []*MerkleNode, data Data) (*MerkleNode, error) {
	nNode := &MerkleNode{children, nil, data}

	// if it is to create a leaf node
	isleaf, err := nNode.IsLeafNode();
	if err != nil {return nil, err}
	if isleaf {
		hash, err := nNode.Data.CalculateHash()
		if err != nil {
			return nil, err
		}
		nNode.Hash = hash[:]
		return nNode, nil
	}
	
	// compute node hash for the intermediate merkle tree node
	spliceHash := []byte{}
	for _, node := range children {
		spliceHash = append(spliceHash, node.Hash...)
	}
	nHash := sha256.Sum256(spliceHash)
	nNode.Hash = nHash[:]

	return nNode, nil
}

// Check is the node is an valid leaf node
//
// Return true if it is leaf node, false if non-leaf node; return non-nil error if it's not
// valid in corresponding node type
func (mn *MerkleNode) IsLeafNode() (bool, error) {
	if mn.Children == nil || len(mn.Children) == 0 {
		if mn.Data == nil {
			return true, fmt.Errorf("Data should be provided for leaf nodes ")
		}
		return true, nil
	}

	// it is not allowed to create non-leaf node through non-nil data
	if mn.Data != nil {
		return false, fmt.Errorf("Data should be null for intermediate nodes ")
	}
	return false, nil
}

func (mn *MerkleNode) VerifyNode() ([]byte, error) {
	// if is leaf node
	isleaf, err := mn.IsLeafNode();
	if err != nil {return nil, err}

	if isleaf {
		return mn.Data.CalculateHash()
	}

	if mn.Hash == nil {
		return nil, fmt.Errorf("Hash should not be empty for intermediate nodes ")
	}

	// compute node hash for the intermediate merkle tree node
	spliceHash := []byte{}
	for _, node := range mn.Children {
		hash, err := node.VerifyNode()
		if err != nil {
			return nil, err
		}
		spliceHash = append(spliceHash, hash...)
	}
	nHash := sha256.Sum256(spliceHash)
	if !bytes.Equal(nHash[:], mn.Hash) {
		return nil, fmt.Errorf("Node hash is ")
	}
	return mn.Hash, nil
}
