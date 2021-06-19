# Go-Merkle: An implementation of merkle tree in Golang

# Introduction

This is an empirical implementation of [Merkle Tree](https://en.wikipedia.org/wiki/Merkle_tree). It's target to 
skcrach the high extendable, readable coding framework of the Merkle Tree.

# Design Highlights
## Extendability
Theorectically, Merkle tree could be adapted to different data types, 
tree structures, and hash functions, etc. Therefore, we adopt follow design to make
our implementation extendable.

+ **Data type extension friendly**

    We design the `Data` interface. Any type of content implementing the `Data` interface could constructure a merkle tree
    with our implementation.
    ```go
    type Data interface {
        CalculateHash() ([]byte, error)
        Equals([]byte, error)
    }
    ```

+ **Tree structure extension friendly**

    We design the `Children` slice in node structure to put the children of the node,  
    so that does not limit the number of children.

    + Binary Merkle Tree: put left and right child of the node in `Children` successively.
    + Multiway Merkle Tree: put all children of the node in `Children` successively.


    ```go
    type MerkleNode struct {
        Children []*MerkleNode // allow Multiway Trees
        Hash     []byte  // allow different hash function
        Data 	 Data
    }
    ```

+ **Hash function extension friendly**

    We design the byte slice `Hash` in node structure to hold the data digest,  
    so that does not limit the digest length.

## Readability
+ Godoc support
    > Require `godoc`, use `go get -v  golang.org/x/tools/cmd/godoc` to install it.

    Developers could read the API document of the whole package
    ```shell
    (RepoRoot)$ godoc -http=:8080
    ```
    then visit `http://localhost:8080/pkg/go-merkle/` in web browser

+ Comments in line

    All key programming reasons have been commented in line

# Get started
## Project Structure
+ `doc.go`: the independent package overview file for godoc
+ `mn.go`: the node-level data structures and APIs
+ `mt.go`: the tree-level data structures and APIs

## Unit Test
+ Run all unit tests
    ```shell
    (RepoRoot)$ go test .
    ```
