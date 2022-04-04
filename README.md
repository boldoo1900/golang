# Go lang

Basic go settings, testing, godoc, profiling, test programms (MACOS)

## Overview
1. [Install prerequisites](#install-prerequisites)

    Golang installation, and environment settings

2. [Basic commands](#basic-commands)

    build, install, run

3. [Golang testing](#golang-testing)

    Test files, syntax, execute

4. [Golang doc](#golang-doc)

    Syntax, execution, beware points

4. [Profiling](#profiling)

    Profiling type, syntax, commands

5. [Other](#other)

    Examples, related web url's

___

## Install prerequisites

### Environment settings (~/.bash_profile)

```sh
export GOPATH=~/go      (or export GOPATH=~/go/dir1:~/dir1:~/dir2)
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
```

### Init mod file
```
go mod init <module_name>
```

## Basic commands
go lang build a binary file, note: In linux environment uses (GOOS=linux)

### build
```sh
go build main.go                             (mac)
or
GOOS=linux GOARCH=amd64 go build main.go     (for linux)
```

### run
```sh
go run main.go
or 
time go run main.go     (execution time measurement)
```

### install
```sh
go install main.go
```
* executable binary file will be created on $GOATH/bin
* created binary file work from everywhere


## Golang testing

### basics
* `*_test.go` execute a file that matches
* `go test q3_test.go` execute command


## Golang doc

### basics
* `godoc -http=:10090`ロカルでドキュメントを実行 (localhost:10090で見れる)
* `go get golang.org/x/tools/cmd/godoc`上記なければこれ実行する
* `godoc -http=:10090 -goroot=$HOME/go`するとdirectory指定できる



## Profiling

Use for speed up program, there is [runtime/pprof](https://golang.org/pkg/runtime/pprof/) and [net/http/pprof](https://golang.org/pkg/net/http/pprof/) type of pprof exists 


### graphical interface package
```sh
brew install graphviz
```

### runtime/pprof - syntax
```sh
import (
    ...
    ...
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    ...
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


    ...
    ...

    if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

}
```

### runtime/pprof - commands
```sh
go build && ./pprof-runtime -cpuprofile cpu.prof -memprofile mem.prof 40       (40 - go program argument)

go tool pprof -http localhost:8888 mem.prof
go tool pprof -http localhost:8888 cpu.prof
```
* See more from fibfib/pprof-runtime.go


### net/http/pprof - syntax 
```sh
import (
    ...
	"log"
	"net/http"
	...
	_ "net/http/pprof"
)

func main(){
    go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

    ...
    ...

}

```

### net/http/pprof - commands
```sh
go build && ./pprof-nethttp 45     or      go run pprof-nethttp.go 45

go tool pprof -http localhost:8888 "http://localhost:6060/debug/pprof/profile?seconds=5"
go tool pprof -http localhost:8888 http://localhost:6060/debug/pprof/heap

```
* See more from fibfib/pprof-nethttp.go



## Other

### related url's

- [ISUCON9](https://github.com/isucon/isucon9-qualify)
- [Mocking HTTP Requests](https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/)
- [Protocol Buffer sample](https://www.grpc.io/docs/quickstart/go/)
- [Build RestAPI](https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da)
