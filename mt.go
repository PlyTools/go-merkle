package go_merkle

import "fmt"

type MerkleTree struct {
	Root *MerkleNode
}

func NewBinaryMerkleTree(dl []Data) (*MerkleTree, error) {
	var nodes []*MerkleNode

	if dl == nil || len(dl) == 0 {
		return nil, fmt.Errorf("data list should not be empty")
	}

	// make the number of leaf nodes even, if it is not before
	if len(dl)%2 != 0 {
		dl = append(dl, dl[len(dl)-1])
	}

	for _, dlitem := range dl {
		node, err := NewMerkleNode(nil, dlitem)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	for i := 0; i < len(dl)/2; i++ {
		var newNodes []*MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			children := []*MerkleNode{}
			children = append(children, nodes[j])
			if j+1 < len(nodes) {
				children = append(children, nodes[j+1])
			}
			node, err := NewMerkleNode(children, nil)
			if err != nil {
				return nil, err
			}
			newNodes = append(newNodes, node)
		}

		nodes = newNodes
	}

	mTree := MerkleTree{nodes[0]}
	return &mTree, nil
}

func (mt *MerkleTree) VerifyTree() ([]byte, error) {
	hash, err := mt.Root.VerifyNode()
	if err != nil {
		return nil, err
	}
	return hash, nil
}
