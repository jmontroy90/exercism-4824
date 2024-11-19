package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

var (
	ErrBufferFull = errors.New("buffer is full")
	ErrReadEmpty  = errors.New("read returned empty byte")
)

type Buffer struct {
	buf     []byte
	readIx  int
	writeIx int
	size    int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		buf: make([]byte, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.size == 0 {
		return 0, ErrReadEmpty
	}
	bs := b.buf[b.readIx]
	b.readIx = (b.readIx + 1) % len(b.buf)
	b.size--
	return bs, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.size == len(b.buf) {
		return ErrBufferFull
	}
	b.buf[b.writeIx] = c
	b.writeIx = (b.writeIx + 1) % len(b.buf)
	b.size++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.size == len(b.buf) {
		_, _ = b.ReadByte() // throw out the next value
	}
	_ = b.WriteByte(c) // this is guaranteed to be in a non-full buffer
}

func (b *Buffer) Reset() {
	// We don't actually re-allocate memory in circular buffers.
	b.writeIx = 0
	b.readIx = 0
	b.size = 0
}
