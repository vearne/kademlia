package kademlia

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BitSet struct {
	Size int
	Data []byte
}

func NewBitSet(size int) *BitSet {
	div, mod := size/8, size%8
	if mod > 0 {
		div++
	}
	return &BitSet{size, make([]byte, div)}
}


func NewBitSetFromRawStr(str string)(*BitSet, error) {
	bs := NewBitSet(len(str) * 8)
	bs.Data = []byte(str)
	return bs, nil
}

func NewBitSetFromBinaryStr(str string) (*BitSet, error) {
	// "00101101"
	var start, end, N int
	N = len(str)
	bs := NewBitSet(N)
	for i := 0; i*8 < N; i++ {
		start = N - (i+1)*8
		if start <= 0 {
			start = 0
		}
		end = N - i*8
		value, err := strconv.ParseUint(str[start:end], 2, 64)
		if err != nil {
			return nil, errors.New("str format error")
		}
		bs.Data[i] = uint8(value)
	}
	return bs, nil
}

func (bs *BitSet) Get(bitIndex int) bool {
	div, mod := bitIndex/8, bitIndex%8
	return (bs.Data[div] >> uint8(mod) & 1) > 0
}

func (bs *BitSet) Set(bitIndex int) {
	div, mod := bitIndex/8, bitIndex%8
	bs.Data[div] |= uint8(1) << uint(mod)
}

func (bs *BitSet) Clear(bitIndex int) {
	div, mod := bitIndex/8, bitIndex%8
	bs.Data[div] &= ^(uint8(1) << uint(mod))
}

func (bs *BitSet) ClearAll() {
	for i := 0; i < len(bs.Data); i++ {
		bs.Data[i] = 0
	}
}

func (bs *BitSet) XOR(other *BitSet) (*BitSet, error) {
	if bs.Size != other.Size {
		return nil, errors.New("BitSet XOR, but size not equal")
	}
	result := NewBitSet(bs.Size)
	for i := 0; i < len(bs.Data); i++ {
		result.Data[i] = bs.Data[i] ^ other.Data[i]
	}
	return result, nil
}

func (bs *BitSet) OR(other *BitSet) (*BitSet, error){
	if bs.Size != other.Size {
		return nil, errors.New("BitSet OR, but size not equal")
	}
	result := NewBitSet(bs.Size)
	for i := 0; i < len(bs.Data); i++ {
		result.Data[i] = bs.Data[i] | other.Data[i]
	}
	return result, nil
}

func (bs *BitSet) AND(other *BitSet) (*BitSet, error){
	if bs.Size != other.Size {
		return nil, errors.New("BitSet OR, but size not equal")
	}
	result := NewBitSet(bs.Size)
	for i := 0; i < len(bs.Data); i++ {
		result.Data[i] = bs.Data[i] & other.Data[i]
	}
	return result, nil
}

func (bs *BitSet) equal(other *BitSet) bool{
	if bs.Size != other.Size{
		return false
	}
	for i := 0; i < len(bs.Data); i++ {
		if bs.Data[i] != other.Data[i]{
			return false
		}
	}
	return true
}

// return binary string
func (bs *BitSet) String() string {
	tempSlice := make([]string, 0)
	div := bs.Size / 8
	for i := 0; i < bs.Size/8; i++ {
		tempSlice = append(tempSlice, fmt.Sprintf("%08b", bs.Data[i]))
	}
	mod := bs.Size % 8
	if mod > 0 {
		tempSlice = append(tempSlice, fmt.Sprintf("%0"+strconv.Itoa(mod)+"b",
			bs.Data[div]))
	}
	reverse(tempSlice)
	return strings.Join(tempSlice, "")
}
