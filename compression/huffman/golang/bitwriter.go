package huffman

import (
	"errors"
	"io"
	"strings"
)

type BitWriter interface {
	WriteOne() error
	WriteZero() error
	FlushBuffer() error
}

type BitWriterImpl struct {
	w        io.Writer
	buff     [1]byte
	bitIndex uint8
}

func NewBitWriter(w io.Writer) *BitWriterImpl {
	return &BitWriterImpl{
		w: w,
	}
}

func (b *BitWriterImpl) WriteOne() error {
	b.buff[0] = b.buff[0] | (1 << (8 - b.bitIndex))
	b.bitIndex++
	if b.bitIndex > 8 {
		return b.FlushBuffer()
	}
	return nil
}

func (b *BitWriterImpl) WriteZero() error {
	b.bitIndex++
	if b.bitIndex > 8 {
		return b.FlushBuffer()
	}
	return nil
}

func (b *BitWriterImpl) FlushBuffer() error {
	n, err := b.w.Write(b.buff[:])
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("did not write byte")
	}
	b.buff[0] = 0
	b.bitIndex = 0
	return nil
}

type BinaryStringWriter struct {
	buffer strings.Builder
}

func (w *BinaryStringWriter) WriteOne() error {
	w.buffer.WriteRune('1')
	return nil
}

func (w *BinaryStringWriter) WriteZero() error {
	w.buffer.WriteRune('0')
	return nil
}

func (w *BinaryStringWriter) FlushBuffer() error {
	return nil
}
