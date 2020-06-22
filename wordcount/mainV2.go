package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

const maxWordLength int = 13

func main() {
	flag.Parse()

	fileName := flag.Args()[0]
	counter := readAndCount(fileName)
	printTop(counter, 3)
}

func readAndCount(fileName string) map[uint64]int {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	counter := map[uint64]int{} // store value by uint64 (largest positive value type in golang)
	b := bufio.NewReader(fp)
	for {
		line, err := b.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		// prepare predefined slice
		template := [maxWordLength]uint64{1}
		for i := 1; i < maxWordLength; i++ {
			template[i] = 26 * template[i-1]
		}

		var wordlen int = 0
		var wordhash uint64 = 0
		for _, b := range line {
			var tmp uint64 = byteToInt(b)
			if tmp != 0 && wordlen < maxWordLength {
				wordhash += tmp * template[wordlen] // z=26  z*1+z*26+z*676
				wordlen++
			} else if 0 < wordlen {
				counter[wordhash]++
				wordhash = tmp // reset hash

				if wordlen >= maxWordLength {
					fmt.Println(wordlen)
				}

				// reset word length
				if tmp == 0 {
					wordlen = 0
				} else {
					wordlen = 1
				}
			}
		}
	}

	return counter
}

func byteToInt(chara byte) uint64 {
	if 'A' <= chara && chara <= 'Z' {
		return uint64(chara-'A') + 1
	} else if 'a' <= chara && chara <= 'z' {
		return uint64(chara-'a') + 1
	} else {
		return 0
	}
}

// uint64 range: 0 - 18446744073709551615
// which mean able to hold characters up to 13
// byteToInt and uint64ToString both reverse process
func uint64ToString(num uint64) string {
	S := ""
	for 0 < num {
		num--
		S += string('a' + (num % 26))
		num /= 26
	}
	return S
}

func printTop(counter map[uint64]int, n int) {
	var kvs []struct {
		key   uint64
		value int
	}
	for k, v := range counter {
		kvs = append(kvs, struct {
			key   uint64
			value int
		}{key: k, value: v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})

	for i := 0; i < n && i < len(kvs); i++ {
		fmt.Println(uint64ToString(kvs[i].key), kvs[i].value)
	}
}

// func byteArrayToString(bs []byte) string {
// 	return *(*string)(unsafe.Pointer(&bs))
// }
