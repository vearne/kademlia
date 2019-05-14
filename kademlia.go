package kademlia


type Node struct {
	ID   BitSet
	IP   string
	// UDP
	port int
}
type KBucket struct {
	NodeSlice []Node
}

func (kb *KBucket) Size() int{
	return len(kb.NodeSlice)
}

type Kademlia struct {
	// inner的key为子树与当前节点的最长公共前缀
	inner map[string]KBucket
	// k桶中元素的数量
	k int
	id BitSet
}

func NewKademlia(id *BitSet, k int) *Kademlia{
	result := &Kademlia{inner:make(map[string]KBucket, id.Size), k:k}

	return result
}

func (k *Kademlia) Store(node *Node){
	//k.id.XOR(node.ID)
}

