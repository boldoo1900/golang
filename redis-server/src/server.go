package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"server/utils"
	"strings"
)

var count = 0

func handleConnection(conn net.Conn, clientID int) {
	for {
		byteArr, err := utils.NewReader(conn).ReadObject()
		// fmt.Println(string(byteArr))
		if err != nil {
			fmt.Println(err)
			return
		}

		tmp := strings.TrimSpace(string(byteArr))
		result := utils.DoLogic(tmp)
		utils.NewWriter(conn).WriteString(result)
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := "0.0.0.0:" + arguments[1]
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
