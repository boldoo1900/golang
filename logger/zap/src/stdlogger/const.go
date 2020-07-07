package stdlogger

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
)

type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
}

var Config = Configuration{
	EnableConsole:     true,
	ConsoleJSONFormat: false,
	ConsoleLevel:      "debug",
}
