package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	flag.Parse()
	fileName := flag.Args()[0]
	counter := readAndCount(fileName)
	printTop(counter, 3)
}

func isAlphabet(c byte) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.Match([]byte{c})
}

func readAndCount(fileName string) map[string]int {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	counter := map[string]int{}
	word := ""
	buf := [1]byte{}
	for {
		n, err := fp.Read(buf[:])
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		c := buf[0]
		if isAlphabet(c) {
			word += string(c)
		} else if word != "" {
			counter[word]++
			word = ""
		}
	}

	return counter
}

func printTop(counter map[string]int, n int) {
	var kvs []struct {
		key   string
		value int
	}
	for k, v := range counter {
		kvs = append(kvs, struct {
			key   string
			value int
		}{key: k, value: v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})

	for i := 0; i < n && i < len(kvs); i++ {
		fmt.Println(kvs[i].key, kvs[i].value)
	}
}
