package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

var array = []int{} // 特に意味のないslice

func fib(n int) int {
	if n <= 1 {
		return n
	}
	res := fib(n-1) + fib(n-2)
	array = append(array, res) // 特に意味のないappend
	return res
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	flag.Parse()
	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(fib(n))
}
