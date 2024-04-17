package bytey

import (
	"errors"
	"fmt"
	"io"
)

// Buffer implements io.ReadWriteSeeker.
type Buffer struct {
	buf    []byte
	offset int64
}

var _ io.ReadWriteSeeker = &Buffer{}

// NewBuffer returns a new *Buffer that implements io.ReadWriteSeeker.
func NewBuffer(buf []byte) *Buffer {
	if buf == nil {
		buf = make([]byte, 0, 100)
	}

	return &Buffer{
		buf:    buf,
		offset: 0,
	}
}

func (b *Buffer) Bytes() []byte {
	return b.buf
}

func (b *Buffer) Len() int {
	return len(b.buf)
}

func (b *Buffer) Read(buf []byte) (int, error) {
	available := len(b.buf) - int(b.offset)
	if available == 0 {
		return 0, io.EOF
	}

	size := len(buf)
	if size > available {
		size = available
	}

	copy(buf, b.buf[b.offset:b.offset+int64(size)])
	b.offset += int64(size)
	return size, nil
}

func (b *Buffer) Write(buf []byte) (int, error) {
	copied := copy(b.buf[b.offset:], buf)
	if copied < len(buf) {
		b.buf = append(b.buf, buf[copied:]...)
	}

	b.offset += int64(len(buf))
	return len(buf), nil
}

var (
	ErrOutOfBounds   = errors.New("out of bounds")
	ErrInvalidWhence = errors.New("invalid whence")
)

func (fb *Buffer) Seek(offset int64, whence int) (int64, error) {
	var newOffset int64
	switch whence {
	case io.SeekStart:
		newOffset = offset
	case io.SeekCurrent:
		newOffset = fb.offset + offset
	case io.SeekEnd:
		newOffset = int64(len(fb.buf)) + offset
	default:
		return 0, fmt.Errorf("whence %d: %w", whence, ErrInvalidWhence)
	}
	if newOffset > int64(len(fb.buf)) || newOffset < 0 {
		return 0, fmt.Errorf("offset %d: %w", offset, ErrOutOfBounds)
	}
	fb.offset = newOffset
	return newOffset, nil
}
