// Disclaimer: I have no idea if this is the correct approach to do a logger, but it seems to work so it's ok
//
// Example usage:
//
//		func main() {
//		    logger.SetLogLevel(logger.DebugLogLevel)
//		    logger.LogDebug("Hello from debug")
//	     	logger.Log("Hi")
//		}
package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type LogLevel int

const (
	NormalLogLevel LogLevel = iota
	InfoLogLevel
	DebugLogLevel
)

type loggerFunc func(format string, args ...any)
type loggerType struct {
	level LogLevel
}

var logger *loggerType = &loggerType{level: NormalLogLevel} // singleton!

func SetLogLevel(level LogLevel) {
	logger = &loggerType{
		level: level,
	}
}

// Decorator for DEBUG/WARN/FATAL messages to print file and line before message
func logWithStack(next loggerFunc) loggerFunc {
	return func(format string, args ...interface{}) {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			next(format, args...)
			return
		}
		loc := fmt.Sprintf("%s:%d", filepath.Base(file), line)
		formatStack := fmt.Sprintf("%-15s | %s", loc, format) // just make sure the '15' number is bigger than any .go file name
		next(formatStack, args...)
	}
}

func logDebug(format string, args ...any) {
	if logger.level == DebugLogLevel {
		log.Printf("[DEBUG]: "+format, args...)
	}
}

func logWarn(format string, args ...any) {
	if logger.level >= InfoLogLevel {
		log.Printf("[WARN]:  "+format, args...)
	}
}

func logInfo(format string, args ...any) {
	if logger.level >= InfoLogLevel {
		log.Printf("[INFO]: "+format, args...)
	}
}

func logNormal(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func logFatal(err error) {
	log.Fatal("[FATAL]: " + err.Error())
}

var LogDebug loggerFunc = logWithStack(logDebug)
var LogWarn loggerFunc = logWithStack(logWarn)
var LogInfo loggerFunc = logInfo
var Log loggerFunc = logNormal
var LogFatal = logFatal
