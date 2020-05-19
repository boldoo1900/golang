package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

const (
	fSimpleString = '+'
	fBulkString   = '$'
	fInteger      = ':'
	fArray        = '*'
	fError        = '-'
)

var (
	errInvalidSyntax = errors.New("resp: invalid syntax")
)

// RedisReadObject read settings
func RedisReadObject(r *bufio.Reader) ([]byte, error) {
	line, err := r.ReadBytes('\n')

	if err != nil {
		return nil, err
	}

	if len(line) <= 0 || line[len(line)-2] != '\r' {
		return nil, errInvalidSyntax
	}

	switch line[0] {
	case fSimpleString, fInteger, fError:
		return line, nil
	case fBulkString:
		return readBulkString(r, line)
	case fArray:
		return readarray(r, line)
	default:
		return nil, errInvalidSyntax
	}
}

func readBulkString(r *bufio.Reader, line []byte) ([]byte, error) {
	count, err := getCount(line)
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

func readarray(r *bufio.Reader, line []byte) ([]byte, error) {
	// Get number of array elements.
	count, err := getCount(line)
	if err != nil {
		return nil, err
	}

	// Read `count` number of RESP objects in the array.
	for i := 0; i < count; i++ {
		buf, err := RedisReadObject(r)
		if err != nil {
			return nil, err
		}
		line = append(line, buf...)
	}

	return line, nil
}

func getCount(line []byte) (int, error) {
	end := bytes.IndexByte(line, '\r')
	return strconv.Atoi(string(line[1:end]))
}
