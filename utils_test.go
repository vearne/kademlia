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

func TestReverse(t *testing.T) {
	slice := []string{"1", "2", "3"}
	reverse(slice)
	if sliceEqual(slice, []string{"3", "2", "1"}) {
		t.Logf("success, %v", true)
	} else {
		t.Errorf("error, expect:%v, got:%v\n", true, false)
	}
}

func TestReverse2(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	reverse(slice)
	if sliceEqual(slice, []string{"4", "3", "2", "1"}) {
		t.Logf("success, %v", true)
	} else {
		t.Errorf("error, expect:%v, got:%v\n", true, false)
	}
}



