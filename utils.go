package kademlia

import (
	"fmt"
	"strings"
)

func stringToBin(s string) (binString string) {
	ss := make([]string, 10)
	for _, c := range s {
		ss = append(ss, fmt.Sprintf("%08b", c))
	}
	return strings.Join(ss, "")
}

func reverse(slice []string){
	for i:=0;i<len(slice)/2;i++{
		slice[i], slice[len(slice) -1 -i] = slice[len(slice) -1 -i], slice[i]
	}
}

func sliceEqual(sc1 []string, sc2 []string) bool{
	if len(sc1) != len(sc2){
		return false
	}
	for i:=0;i<len(sc1);i++{
		if sc1[i] != sc2[i]{
			return false
		}
	}
	return true
}