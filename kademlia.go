package kademlia

type Node struct {
	ID   string
	IP   string
	port int
}
type KBucket struct {
	NodeSlice []Node
}

type Kademlia struct {
	inner map[string]KBucket
}
