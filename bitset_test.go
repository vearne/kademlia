package kademlia

import (
	"testing"
)

func TestBitSet1(t *testing.T){
	bs := NewBitSet(10)
	bs.Set(2)
	bs.Set(9)
	target := "1000000100"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestBitSet2(t *testing.T){
	bs := NewBitSet(10)
	bs.Set(2)
	bs.Set(8)
	target := "0100000100"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestBitSet3(t *testing.T){
	bs := NewBitSet(8)
	bs.Set(2)
	bs.Set(5)
	target := "00100100"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestBitSet4(t *testing.T){
	bs := NewBitSet(5)
	bs.Set(2)
	target := "00100"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestBitSet5(t *testing.T){
	bs := NewBitSet(5)
	bs.Set(2)
	target := "00100"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestNewBitSetFromStr(t *testing.T){
	bs, _ := NewBitSetFromBinaryStr("00110101")
	target := "00110101"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestNewBitSetFromStr2(t *testing.T){
	bs, _ := NewBitSetFromBinaryStr("0101")
	target := "0101"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestNewBitSetFromStr3(t *testing.T){
	bs, _ := NewBitSetFromBinaryStr("010111000011")
	target := "010111000011"
	result := bs.String()
	if bs.String() == target{
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestGet(t *testing.T){
	bs, _ := NewBitSetFromBinaryStr("010111000011")
	result := bs.Get(6)
	target := true
	if result {
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}

func TestGet2(t *testing.T){
	bs, _ := NewBitSetFromBinaryStr("010111000011")
	result := bs.Get(5)
	target := true
	if !result {
		t.Logf("success, %v", bs.String())
	} else {
		t.Errorf("error, expect:%v, got:%v\n", target, result)
	}
}