package kademlia

import (
	"testing"
)

func TestStringToBin(t *testing.T) {
	x := stringToBin("cc")

	if x == "0110001101100011" {
		t.Logf("success, %v", x)
	} else {
		t.Errorf("error, expect:%v, got:%v\n", "0110001101100011", x)
	}
}
