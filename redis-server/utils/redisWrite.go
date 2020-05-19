package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	prefString  = []byte{'+'}
	prefInteger = []byte{':'}
	prefarray   = []byte{'*'}
	prefBulk    = []byte{'$'}
	prefError   = []byte{'-'}
	prefCRLF    = []byte{'\r', '\n'}
)

var mapObj = make(map[string]string)
var mLock = &sync.Mutex{}

func stringToByte(str string) []byte {
	return []byte(strconv.Quote(str))
}

func rSet(key string, value string) {
	mLock.Lock()
	mapObj[key] = value
	mLock.Unlock()
}

func rGet(key string) (string, bool) {
	mLock.Lock()
	defer mLock.Unlock()

	value, hasError := mapObj[key]
	return value, hasError
}

func rDelete(key string) {
	mLock.Lock()
	defer mLock.Unlock()
	delete(mapObj, key)
}

func rIncrement(key string, inc string) (string, bool) {
	mLock.Lock()
	defer mLock.Unlock()

	value, hasVal := mapObj[key]
	if hasVal {
		oldVal, err := strconv.Atoi(value)
		if err != nil {
			return "value is not an integer or out of range", false
		}

		incVal, err := strconv.Atoi(inc)
		if err != nil {
			return "value is not an integer or out of range", false
		}

		mapObj[key] = strconv.Itoa(oldVal + incVal)
		return mapObj[key], true
	}

	_, err := strconv.Atoi(inc)
	if err != nil {
		return "value is not an integer or out of range", false
	}

	mapObj[key] = inc
	return inc, true
}

// RedisWriteObject response
func RedisWriteObject(w *bufio.Writer, commandStr string) {
	// fmt.Println(commandStr)
	strArr := strings.Split(commandStr, "\r\n")
	if len(strArr) < 3 {
		w.Write(prefString)
		w.Write([]byte("command not found"))
		w.Write(prefCRLF)
		w.Flush()
		return
	}

	fmt.Println(len(strArr), ": ", strArr)
	switch strings.ToUpper(strArr[2]) {
	case "PING":
		{
			w.Write(prefString)
			if len(strArr) == 5 {
				w.Write(stringToByte(strArr[4]))
			} else if len(strArr) <= 4 {
				w.WriteString("PONG")
			} else {
				w.WriteString("wrong number of arguments for 'ping' command")
			}
			w.Write(prefCRLF)
		}
	case "GET":
		{
			w.Write(prefBulk)
			if len(strArr) == 5 {
				// mapObj[strArr[4]]
				if val, ok := rGet(strArr[4]); ok {
					w.WriteString(strconv.Itoa(len(val)))
					w.Write(prefCRLF)
					w.WriteString(val)
				}

				w.WriteString("-1")
			} else {
				msg := "wrong number of arguments for 'get' command"

				w.WriteString(strconv.Itoa(len(msg)))
				w.Write(prefCRLF)
				w.WriteString(msg)
			}
			w.Write(prefCRLF)
		}
	case "SET":
		{
			if len(strArr) >= 7 && len(strArr) <= 9 {
				// 7: val, 5: key

				if len(strArr) == 9 { // nx, xx option
					if strings.ToUpper(strArr[8]) == "NX" {
						// mapObj[strArr[4]]
						if _, ok := rGet(strArr[4]); ok {
							w.Write(prefBulk)
							w.WriteString("-1")
						} else {
							w.Write(prefString)
							rSet(strArr[4], strArr[6])
							w.WriteString("OK")
						}
					} else if strings.ToUpper(strArr[8]) == "XX" {
						// mapObj[strArr[4]]
						if _, ok := rGet(strArr[4]); ok {
							w.Write(prefString)
							rSet(strArr[4], strArr[6])
							w.WriteString("OK")
						} else {
							w.Write(prefBulk)
							w.WriteString("-1")
						}
					}
				} else {
					w.Write(prefString)
					rSet(strArr[4], strArr[6])
					w.Write(stringToByte("OK"))
				}
			} else {
				msg := "wrong number of arguments for 'set' command"
				w.WriteString(msg)
			}
			w.Write(prefCRLF)
		}
	case "DEL":
		{
			if len(strArr) >= 4 {
				w.Write(prefInteger)
				delNum := 0
				for i := 4; i < len(strArr); i++ { // i+2 not working, investigation required
					if _, ok := rGet(strArr[i]); ok {
						rDelete(strArr[i])
						// delete(mapObj, strArr[i])
						delNum++
					}
				}
				w.WriteString(strconv.Itoa(delNum))
			} else {
				w.Write(prefString)
				msg := "wrong number of arguments for 'del' command"
				w.WriteString(msg)
			}

			w.Write(prefCRLF)
		}
	case "INCRBY":
		{
			if len(strArr) == 7 {
				result, isSuccess := rIncrement(strArr[4], strArr[6])
				if isSuccess {
					w.Write(prefInteger)
					w.WriteString(result)
				} else {
					w.Write(prefString)
					w.WriteString(result)
				}
			} else {
				w.Write(prefString)
				w.WriteString("wrong number of arguments for 'incrby' command")
			}
			w.Write(prefCRLF)
		}
	default:
		{
			w.Write(prefString)
			w.Write([]byte("command not found"))
			w.Write(prefCRLF)
		}
	}

	w.Flush()
}
