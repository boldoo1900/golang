package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sort"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	fileName := flag.Args()[0]

	// start := time.Now()
	counter := readAndCount(fileName)
	// elapsed := time.Since(start)
	// fmt.Println(elapsed)
	// fmt.Println(counter)
	printTop(counter, 3)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

var IsLetter = regexp.MustCompile(`^[a-z]+$`).MatchString

func readAndCount(fileName string) map[string]int {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	counts := map[string]int{}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		if IsLetter(word) {
			counts[word]++
		}
	}

	return counts
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
