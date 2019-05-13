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