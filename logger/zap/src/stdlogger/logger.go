package stdlogger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})

	Errorf(args ...interface{})

	WithFields(keyValues Fields) Logger
}

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func New() (Logger, error) {
	cores := []zapcore.Core{}

	if Config.EnableConsole {
		level := getZapLevel(Config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(Config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	// if config.EnableFile {
	// 	level := getZapLevel(config.FileLevel)
	// 	writer := zapcore.AddSync(&lumberjack.Logger{
	// 		Filename: config.FileLocation,
	// 		MaxSize:  100,
	// 		Compress: true,
	// 		MaxAge:   28,
	// 	})
	// 	core := zapcore.NewCore(getEncoder(config.FileJSONFormat), writer, level)
	// 	cores = append(cores, core)
	// }

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	// logger, _ := zap.NewDevelopment()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debug(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Info(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warning(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Error(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Errorf(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatal(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
