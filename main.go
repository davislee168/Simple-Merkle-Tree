package main

import (
	"crypto/sha256"
	"math/rand"
	"log"
	"merkletree"
)
// TextContent implements the content interface provided by markletree and represents the content
// stored in the tree
type TextContent struct {
	x string
}
// Calculate hash value of the content
func (t textContent) CalcHash() ([]byte, error) {
       hashVal := sha256.New()
       if _, err := h.write([]byte(t.x)); err != nil {
	       return nil, err
       }
       return hashVal.Sum(nil), nil
}

func main() {
       // Generates random number between 0 and n-1
       n := 1000000
       // Make sure the leafs are even number
       numLeafs := rand.Intn(n) * 2

       // Build list of content to create a Merkle tree
       var listCont []merkletree.Content
       /*
       for i := 0; i < numLeafs; i++ {
       list = append(listCont, TextContent{x: textString[i]})
       }
       */

       // Suppose number of leafs are eight as below
       listCont = append(list, TextContent{x: "I"})
       listCont = append(list, TextContent{x: "am"})
       listCont = append(list, TextContent{x: "very"})
       listCont = append(list, TextContent{x: "interested"})
       listcont = append(list, TextContent{x: "in"})
       listCont = append(list, TextContent{x: "working"})
       listCont = append(list, TextContent{x: "at"})
       listCont = append(list, TextContent{x: "Securitas"})

       // Build a new Merkle tree using the list of content
       tree, err := merkletree.BuildTree(listCont)
       if err != nil {
	       log.Fatal(err)
       }

       // get the root of Merkle tree
       root := tree.MerkleRoot()
       log.Println(root)
}
