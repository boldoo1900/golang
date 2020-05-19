package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	redis "gopkg.in/redis.v5"
)

func main() {
	host := flag.String("host", "localhost", "servert host")
	port := flag.Int("port", 6379, "servert port")
	parallelism := flag.Int("par", 10, "parallel clients")
	repeat := flag.Int("rep", 10, "benchmark repeat times")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	cli := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	defer cli.Close()

	var result bool

	result = testCommand(cli)
	if !result {
		println("command test failed")
		return
	}
	println("command test succeed!")

	testBench(cli, *parallelism, *repeat)
}

func testCommand(cli *redis.Client) bool {
	var (
		status  *redis.StatusCmd
		strcmd  *redis.StringCmd
		boolcmd *redis.BoolCmd
		intcmd  *redis.IntCmd
	)

	// PING
	status = cli.Ping()
	if res, err := status.Result(); err != nil {
		fmt.Printf("PING Error: %v\n", err)
		return false
	} else if res != "PONG" {
		fmt.Printf("PING Error: response is %s, not PONG\n", res)
		return false
	}

	// GET, SET, EXISTS, DEL

	var (
		testKey    = uuid.Must(uuid.NewV4()).String()
		testValue  = uuid.Must(uuid.NewV4()).String()
		testValue2 = uuid.Must(uuid.NewV4()).String()
	)

	/*
		boolcmd = cli.Exists(testKey)
		if exist, err := boolcmd.Result(); err != nil {
			fmt.Printf("EXISTS Error: %v\n", err)
			return false
		} else if exist {
			fmt.Printf("Key %s must not be exists\n", testKey)
			return false
		}
	*/

	status = cli.Set(testKey, testValue, 0)
	if err := status.Err(); err != nil {
		fmt.Printf("SET Error: %v\n", err)
		return false
	}

	/*
		boolcmd = cli.Exists(testKey)
		if exist, err := boolcmd.Result(); err != nil {
			fmt.Printf("EXISTS Error: %v\n", err)
			return false
		} else if !exist {
			fmt.Printf("Key %s must be exists\n", testKey)
			return false
		}
	*/

	strcmd = cli.Get(testKey)
	if val, err := strcmd.Result(); err != nil {
		fmt.Printf("GET Error: %v\n", err)
		return false
	} else if val != testValue {
		fmt.Printf("GET Error: wrong value, expected %s actually %s\n", testValue, val)
		return false
	}

	boolcmd = redis.NewBoolCmd("set", testKey, testValue2, "nx")
	cli.Process(boolcmd)
	if ok, err := boolcmd.Result(); err != nil {
		fmt.Printf("SET NX Error: %v\n", err)
		return false
	} else if ok {
		fmt.Printf("SET NX must not be success for key %s\n", testKey)
		return false
	}

	boolcmd = cli.SetXX(testKey, testValue2, 0)
	if ok, err := boolcmd.Result(); err != nil {
		fmt.Printf("SET XX Error: %v\n", err)
		return false
	} else if !ok {
		fmt.Printf("SET XX must be success for key %s\n", testKey)
		return false
	}

	strcmd = cli.Get(testKey)
	if val, err := strcmd.Result(); err != nil {
		fmt.Printf("GET Error: %v\n", err)
		return false
	} else if val != testValue2 {
		fmt.Printf("Get Error: wrong value, expected %s actually %s\n", testValue2, val)
		return false
	}

	intcmd = cli.Del(testKey)
	if val, err := intcmd.Result(); err != nil {
		fmt.Printf("DEL Error: %v\n", err)
		return false
	} else if val != 1 {
		fmt.Printf("DEL Error: wrong value, expected 1 actually %d\n", val)
		return false
	}

	/*
		boolcmd = cli.Exists(testKey)
		if exist, err := boolcmd.Result(); err != nil {
			fmt.Printf("EXISTS Error: %v\n", err)
			return false
		} else if exist {
			fmt.Printf("Key %s must not be exists\n", testKey)
			return f.alse
		}
	*/

	boolcmd = cli.SetXX(testKey, testValue, 0)
	if ok, err := boolcmd.Result(); err != nil {
		fmt.Printf("SET XX Error: %v\n", err)
		return false
	} else if ok {
		fmt.Printf("SET XX must not be success for key %s\n", testKey)
		return false
	}

	//INCRBY, DECRBY

	boolcmd = redis.NewBoolCmd("set", testKey, 10, "nx")
	cli.Process(boolcmd)
	if ok, err := boolcmd.Result(); err != nil {
		fmt.Printf("SETNX Error: %v\n", err)
		return false
	} else if !ok {
		fmt.Printf("SET NX must be success for key %s\n", testKey)
		return false
	}

	intcmd = cli.IncrBy(testKey, 15)
	if val, err := intcmd.Result(); err != nil {
		fmt.Printf("INCRBY Error: %v\n", err)
		return false
	} else if val != 25 {
		fmt.Printf("INCRBY Error: wrong value, expected 25 actually %d\n", val)
		return false
	}

	/*
		intcmd = cli.DecrBy(testKey, 30)
		if val, err := intcmd.Result(); err != nil {
			fmt.Printf("DECRBY Error: %v\n", err)
			return false
		} else if val != -5 {
			fmt.Printf("DECRBY Error: wrong value, expected -5 actually %d\n", val)
			return false
		}
	*/
	return true
}

func testBench(cli *redis.Client, par, rep int) {
	var wait sync.WaitGroup

	for i := 0; i < par; i += 1 {
		wait.Add(1)
		go func(wait *sync.WaitGroup, seq int) {
			defer wait.Done()
			bench(cli, seq, rep)
		}(&wait, i)
	}

	wait.Wait()
}

func bench(cli *redis.Client, seq, rep int) {
	start := time.Now()

	for r := 0; r < rep; r += 1 {
		var keynum = 100
		keys := make([]string, 100)
		for i := 0; i < keynum; i += 1 {
			keys[i] = uuid.Must(uuid.NewV4()).String()
		}

		for i, k := range keys {
			cli.Process(redis.NewBoolCmd("set", k, i, "nx"))
			r := cli.Get(k)
			if val, err := r.Result(); err != nil {
				fmt.Printf("Bench failed [setnx]: %v\n", err)
				return
			} else if val != strconv.Itoa(i) {
				fmt.Printf("Bench failed [setnx]: wrong value e:%d a:%s\n", i, val)
			}
		}

		for i, k := range keys {
			cli.IncrBy(k, 10)
			r := cli.Get(k)
			if val, err := r.Result(); err != nil {
				fmt.Printf("Bench failed [incrby]: %v\n", err)
				return
			} else if val != strconv.Itoa(i+10) {
				fmt.Printf("Bench failed [incrby]: wrong value e:%d a:%s\n", i+10, val)
			}
		}

		for i, k := range keys {
			cli.SetXX(k, i*2, 0)
			r := cli.Get(k)
			if val, err := r.Result(); err != nil {
				fmt.Printf("Bench failed [decrby]: %v\n", err)
				return
			} else if val != strconv.Itoa(i*2) {
				fmt.Printf("Bench failed [decrby]: wrong value e:%d a:%s\n", i*2, val)
			}
		}

		for _, k := range keys {
			cli.Del(k)
		}
	}
	end := time.Now()
	fmt.Printf("client %d: %f second(s)\n", seq, (end.Sub(start)).Seconds())
}
