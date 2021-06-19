/*
Package go_merkle is an empirical implementation of Merkle Tree, see https://en.wikipedia.org/wiki/Merkle_tree.
It's target to skcrach the high extendable, readable coding framework of the Merkle Tree.

There are some design highlights for extendability. Theorectically, Merkle tree could be adapted to different data types,
tree structures, and hash functions, etc. Therefore, we adopt follow design to make
our implementation extendable.

To make it data type extension friendly, we design the `Data` interface. Any type of
content implementing the `Data` interface could constructure a merkle tree
with our implementation.
	type Data interface {
		CalculateHash() ([]byte, error)
		Equals([]byte, error)
	}

To make it tree structure extension friendly, we design the `Children` slice in node
structure to put the children of the node, so that does not limit the number of children.

	type MerkleNode struct {
		Children []*MerkleNode
		Hash     []byte
		Data 	 Data
	}

To make it hash function extension friendly, we design the byte slice `Hash` in node
structure to hold the data digest, so that does not limit the digest length
*/
package go_merkle
