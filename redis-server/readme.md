
## run redis tester
```sh
% ./redis_tester -h                                            
Usage of ./redis_tester:
  -host string
        servert host (default "localhost")
  -par int
        parallel clients (default 10)
  -port int
        servert port (default 6379)
  -rep int
        benchmark repeat times (default 10)

training test option
./redis_tester -port 123 -rep 200
```

## run server
```
go run server.go 123                (server)

redis-cli -p 123                    (redis client)
```

# Reference
* [TCP client server](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)
* [Redis read write](https://www.redisgreen.net/blog/reading-and-writing-redis-protocol/)
* [Reader & Writer](https://gist.github.com/hyper0x/8f724925c344f896b63c)
* [Golang IO article](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)