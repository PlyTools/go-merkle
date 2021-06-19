package go_merkle

import (
	"crypto/sha256"
	"testing"
)

func TestNewBinaryMerkleTree_Nil(t *testing.T) {
	// Case1: test input check with nil
	if nMT, err := NewBinaryMerkleTree(nil); err == nil {
		t.Error("Nil input expected to get nil with error: ", nMT, err)
	}
}

func TestNewBinaryMerkleTree_EmptyList(t *testing.T) {
	// Case2: test input check with empty data [][]byte
	empInput := []Data{}
	if nMT, err := NewBinaryMerkleTree(empInput); err == nil {
		t.Fatal("Empty input expected to get nil with error: ", nMT, err)
	}
}

func TestNewBinaryMerkleTree_OddList(t *testing.T) {
	// Case3: build tree based on list with odd Data
	var otsdl []Data
	otsdl = append(otsdl, TestStringData{data: "Hello"})
	if _, err := NewBinaryMerkleTree(otsdl); err != nil {
		t.Fatal(err)
	} else {
		// TODO: renqian - add comparision with expected result
	}
}

func TestNewBinaryMerkleTree_EvenList(t *testing.T) {
	// Case4: build tree based on list with even Data
	var etsdl []Data
	etsdl = append(etsdl, TestStringData{data: "Hello"}, TestStringData{data: "World"})
	if _, err := NewBinaryMerkleTree(etsdl); err != nil {
		t.Fatal(err)
	} else {
		// TODO: renqian - add comparision with expected result
	}
}

//TestStringData implements the Data interface and represent the data content of Merkle tree
type TestStringData struct {
	data string
}

func (tsd TestStringData) CalculateHash() ([]byte, error) {
	hash := sha256.Sum256([]byte(tsd.data))
	return hash[:], nil
}

func (tsd TestStringData) Equals(ntsd Data) (bool, error) {
	return tsd.data == ntsd.(TestStringData).data, nil
}

func TestMerkleTree_VerifyTree(t *testing.T) {
	var tsdl []Data
	tsdl = append(tsdl, TestStringData{data: "Hello"})
	tsdl = append(tsdl, TestStringData{data: ","})
	tsdl = append(tsdl, TestStringData{data: "I"})
	tsdl = append(tsdl, TestStringData{data: "am"})
	tsdl = append(tsdl, TestStringData{data: "PlyTools"})

	// Create a new Merkle Tree
	var bmt *MerkleTree
	if nbmt, err := NewBinaryMerkleTree(tsdl); err != nil {
		t.Fatal(err)
	} else {
		bmt = nbmt
	}

	// Verify the entire tree (hashes for each node) is valid
	if _, err := bmt.VerifyTree(); err != nil {
		t.Fatal(err)
	}

}


