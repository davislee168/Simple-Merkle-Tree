1.	x = rand()
2.	list = 2 * x to make sure there is an even leafs (duplicate the last leaf n to leaf n+1 if odd leafs 
provided)
3.	build a Merkle tree bases on the list as input using a hash algorithm hash(x) of the nodes left to 
right children�s hashes.  I choose sha256() encryption 
4.	get the hash of the root node of the tree.

In main.go
CalcHash(): calculate hash value of content by sha256() function
Creates text content to construct a Merkle tree
Build a new tree
Get the root of the Merkle tree
