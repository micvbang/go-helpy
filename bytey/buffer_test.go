package bytey_test

import (
	"errors"
	"io"
	"testing"

	"github.com/micvbang/go-helpy/bytey"
	"github.com/micvbang/go-helpy/slicey"
)

// TestBufferSeekEmpty verifies that bytey.ErrOutOfBounds is returned when
// attempting to seek out of bounds.
func TestBufferSeekEmpty(t *testing.T) {
	whences := []int{
		io.SeekStart,
		io.SeekEnd,
		io.SeekCurrent,
	}

	for _, whence := range whences {
		buf := bytey.NewBuffer(nil)
		n, err := buf.Seek(1, whence)
		if !errors.Is(err, bytey.ErrOutOfBounds) {
			t.Fatalf("expected error '%v', got '%v'", bytey.ErrOutOfBounds, err)
		}
		equal(t, 0, n, "expected offset %d, got %d", 0, n)
	}
}

// TestBufferReadEmpty verifies that io.EOF is returned when attempting to read
// an empty buffer.
func TestBufferReadEmpty(t *testing.T) {
	buf := bytey.NewBuffer(nil)
	n, err := buf.Read(make([]byte, 100))
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected error '%v', got '%v'", io.EOF, err)
	}
	equal(t, 0, n, "expected to read %d, got %d", 0, n)
}

// TestBufferWriteSeekRead verifies that Read can read the bytes that were
// written by Write, after we've Seeked to the beginning of the file.
func TestBufferWriteSeekRead(t *testing.T) {
	buf := bytey.NewBuffer(nil)

	write := []byte("0123456789")
	bufWrite(t, buf, write)

	n, err := buf.Read([]byte{})
	if err != io.EOF {
		t.Fatalf("expected EOF, got %s", err)
	}
	equal(t, n, 0, "expected to read %d bytes, read %d", 0, n)
}

// TestBufferOverwrite verifies that Write and Seek can be used to expand and
// overwrite contents of the existing buffer.
func TestBufferOverwrite(t *testing.T) {
	buf := bytey.NewBuffer(nil)

	write := []byte("0123456789")
	{ // write 10 bytes
		bufWrite(t, buf, write)
	}

	{ // read without seeking (no bytes to read)
		n, err := buf.Read([]byte{})
		if err != io.EOF {
			t.Fatalf("expected EOF, got %s", err)
		}
		equal(t, n, 0, "expected to read %d bytes, read %d", 0, n)
	}

	{ // seek to beginning of file (10 bytes left to read)
		newOffset := bufSeek(t, buf, 0, io.SeekStart)
		equal(t, newOffset, 0, "expected offset %d, got %d", 0, newOffset)

		n, read := bufRead(t, buf, make([]byte, 100), 10)
		if !slicey.Equal(write, read[:n]) {
			t.Fatalf("expected to read %s, but read %s", write, read[:n])
		}
	}

	{ // seek to middle of file, write 10 new bytes
		seekTo1 := int64(5)
		newOffset1 := bufSeek(t, buf, seekTo1, io.SeekStart)
		equal(t, newOffset1, seekTo1, "expected offset %d, got %d", seekTo1, newOffset1)

		bufWrite(t, buf, write)

		seekTo2 := int64(0)
		newOffset2 := bufSeek(t, buf, seekTo2, io.SeekStart)
		equal(t, 0, newOffset2, "expected to seek to offset %d, got %d", seekTo2, newOffset2)

		n, read := bufRead(t, buf, make([]byte, 20), 15)
		if !slicey.Equal([]byte("012340123456789"), read[:n]) {
			t.Fatalf("expected to read %s, but read %s", write, read[:n])
		}
	}

	{ // seek to end of file, write 10 new bytes
		newOffset1 := bufSeek(t, buf, 0, io.SeekEnd)
		equal(t, newOffset1, 15, "expected offset %d, got %d", 15, newOffset1)

		bufWrite(t, buf, write)

		newOffset2 := bufSeek(t, buf, 0, io.SeekStart)
		equal(t, newOffset2, 0, "expected offset %d, got %d", 0, newOffset2)

		n, read := bufRead(t, buf, make([]byte, 100), 25)
		if !slicey.Equal([]byte("0123401234567890123456789"), read[:n]) {
			t.Fatalf("expected to read %s, but read %s", write, read[:n])
		}
	}
}

func bufWrite(t *testing.T, buf *bytey.Buffer, bs []byte) {
	n, err := buf.Write(bs)
	noError(t, err, "expected to be able to write, got error: %s", err)
	equal(t, n, len(bs), "expected to write %d, wrote %d", len(bs), n)

}

func bufRead(t *testing.T, buf *bytey.Buffer, bs []byte, expectedBytes int) (int, []byte) {
	n, err := buf.Read(bs)
	noError(t, err, "expected to be able to read, got error: %s", err)
	equal(t, n, expectedBytes, "expected to read %d, read %d", expectedBytes, n)
	return n, bs
}

func bufSeek(t *testing.T, buf *bytey.Buffer, offset int64, whence int) int64 {
	seek, err := buf.Seek(offset, whence)
	noError(t, err, "expected to be able to seek, got error: %s", err)
	return seek
}

func noError(t *testing.T, err error, msg string, args ...any) {
	if err != nil {
		t.Fatalf(msg, args...)
	}
}

func equal[T comparable](t *testing.T, v1 T, v2 T, msg string, args ...any) {
	if v1 != v2 {
		t.Fatalf(msg, args...)
	}
}
