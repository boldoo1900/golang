package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"server/utils"
	"strings"
)

var count = 0
var colors = [5]string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m"} //Red, Green, Yellow, Blue, Purple

func handleConnection(conn net.Conn, clientID int) {
	buffer := make([]byte, 512)
	buffLen, err := conn.Read(buffer)
	if err != nil {
		log.Println("client left..")
		conn.Close()
		return
	}

	commands := strings.TrimSpace(string(buffer[:buffLen]))

	// colors up to 5
	// Println(colors[clientID%5], temp)

	w := bufio.NewWriter(conn)
	utils.RedisWriteObject(w, commands)

	// recursive func to handle io.EOF for random disconnect
	handleConnection(conn, clientID)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := "127.0.0.1:" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, count)
		count++
	}
}
