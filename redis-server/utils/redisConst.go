package utils

const (
	prefString  = "+"
	prefInteger = ":"
	prefarray   = "*"
	prefBulk    = "$"
	prefError   = "-"
	prefCRLF    = "\r\n"
)

const (
	prefByteString  = byte('+')
	prefByteInteger = byte(':')
	prefByteError   = byte('-')
)

const (
	msgErrorDefault = prefString + "command not found"
	msgOK           = prefString + "OK"
	msgNil          = prefBulk + "-1"
)

const seperator = "\r\n"

// Colors print fmt.Println with color
var Colors = [5]string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m"} //Red, Green, Yellow, Blue, Purple
var commandList = []string{"PING", "SET", "GET", "DEL", "INCRBY"}
