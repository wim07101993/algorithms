package huffman

import (
	"io"
)

type BitReader interface {
	Read() (bool, error)
}

type BinaryStringReader struct {
	value string
	index int
}

func NewBinaryStringReader(value string) *BinaryStringReader {
	return &BinaryStringReader{value: value}
}

func (r *BinaryStringReader) Read() (bool, error) {
	if r.index < len(r.value) {
		val := r.value[r.index]
		r.index++
		return val == '1', nil
	}
	return false, io.EOF
}
