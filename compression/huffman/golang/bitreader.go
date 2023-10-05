package huffman

import (
	"errors"
	"io"
)

type BitReader interface {
	Read() (bool, error)
}

type BitReaderImpl struct {
	r        io.Reader
	buff     [1]byte
	bitIndex int8
}

func NewBitReader(r io.Reader) *BitReaderImpl {
	return &BitReaderImpl{
		r:        r,
		bitIndex: -1,
	}
}

func (b *BitReaderImpl) Read() (bool, error) {
	if b.bitIndex < 0 {
		n, err := b.r.Read(b.buff[:])
		if err != nil {
			return false, err
		}
		if n != 1 {
			return false, errors.New("did not read byte")
		}
		b.bitIndex = 8
	}
	value := b.buff[0]&(1<<b.bitIndex) > 1
	b.bitIndex--
	return value, nil
}
