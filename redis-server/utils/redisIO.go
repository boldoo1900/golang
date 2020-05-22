package utils

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// Reader ...
type Reader struct {
	*bufio.Reader
}

// Writer ...
type Writer struct {
	wd *bufio.Writer
}

// NewReader ...
func NewReader(reader io.Reader) *Reader {
	return &Reader{
		Reader: bufio.NewReaderSize(reader, 1024),
	}
}

// NewWriter ...
func NewWriter(wd io.Writer) *Writer {
	return &Writer{
		wd: bufio.NewWriter(wd),
	}
}

// Reset ...
func (r *Reader) Reset(rd io.Reader) {
	r.Reset(rd)
}

// ReadObject ...
func (r *Reader) ReadObject() ([]byte, error) {
	line, err := r.ReadBytes('\n') // read first number
	if err != nil {
		return nil, err
	}

	if line[0] == prefByteString || line[0] == prefByteInteger || line[0] == prefByteError {
		return line, err
	}

	// Get number of array elements.
	count, err := r.getCount(line)
	if err != nil {
		return nil, err
	}

	// Read `count` number of RESP objects in the array.
	for i := 0; i < count; i++ {
		tmpLine, _ := r.ReadBytes('\n')
		buf, err := r.readBulkString(tmpLine) // little bit faster than using ReadSlice & ReadByte 2 times
		if err != nil {
			return nil, err
		}
		line = append(line, buf...)
	}

	return line, nil
}

func (r *Reader) readBulkString(line []byte) ([]byte, error) {
	count, err := r.getCount(line)
	if err != nil {
		return nil, err
	}
	if count == -1 {
		return line, nil
	}

	buf := make([]byte, len(line)+count+2)
	copy(buf, line)
	_, err = io.ReadFull(r, buf[len(line):])
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *Reader) getCount(line []byte) (int, error) {
	end := bytes.IndexByte(line, '\r')
	return strconv.Atoi(string(line[1:end]))
}

// WriteString ...
func (w *Writer) WriteString(content string) error {
	_, err := w.wd.WriteString(content)
	if err == nil {
		err = w.wd.Flush()
	}
	return err
}
