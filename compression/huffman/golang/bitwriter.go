package huffman

import (
	"strings"
)

type BitWriter interface {
	WriteOne() error
	WriteZero() error
	FlushBuffer() error
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
