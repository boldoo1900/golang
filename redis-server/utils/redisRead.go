package utils

import (
	"bufio"
	"bytes"
	"io"
	"net"
)

// Read chat info
func Read(conn net.Conn, delim byte) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}
