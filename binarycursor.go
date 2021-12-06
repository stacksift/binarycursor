package binarycursor

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrReadWrongSize = errors.New("Wrong size")
var ErrReaderInvalid = errors.New("Reader is nil")

type BinaryCursor struct {
	br BinaryReader
}

func NewBinaryCursor(r io.Reader) BinaryCursor {
	return BinaryCursor{
		br: NewBinaryReader(r),
	}
}

func NewBinaryCursorAt(r io.ReaderAt, pos int64) BinaryCursor {
	pr := NewPositionReaderAt(r, pos)

	return NewBinaryCursor(&pr)
}

func (c *BinaryCursor) Read(p []byte) (n int, err error) {
	return c.br.Read(p)
}

func (c *BinaryCursor) Order() binary.ByteOrder {
	return c.br.Order
}

func (c *BinaryCursor) FlipOrder() {
	c.br.FlipOrder()
}

func (c *BinaryCursor) ReadUint8() (uint8, error) {
	return c.br.ReadUint8()
}

func (c *BinaryCursor) ReadUint16() (uint16, error) {
	return c.br.ReadUint16()
}

func (c *BinaryCursor) ReadUint32() (uint32, error) {
	return c.br.ReadUint32()
}

func (c *BinaryCursor) ReadUint64() (uint64, error) {
	return c.br.ReadUint64()
}

func (c *BinaryCursor) ReadNullTerminatedUTF8String() (string, error) {
	return c.br.ReadNullTerminatedUTF8String()
}
