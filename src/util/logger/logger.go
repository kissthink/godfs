package logger

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"util/timeutil"
)

const (
	tracePrefix = "TRACE - " // 0
	debugPrefix = "DEBUG - " // 1
	infoPrefix  = "INFO  - " // 2
	warnPrefix  = "WARN  - " // 3
	errorPrefix = "ERROR - " // 4
	fatalPrefix = "FATAL - " // 5
)

func init() {
	now := time.Now()
	lastLogTime = &now
}

func Trace(o ...interface{}) {
	if logLevel > 0 {
		return
	}
	write(tracePrefix, o)
}

func Debug(o ...interface{}) {
	if logLevel > 1 {
		return
	}
	write(debugPrefix, o)
}

func Info(o ...interface{}) {
	if logLevel > 2 {
		return
	}
	write(infoPrefix, o)
}

func Warn(o ...interface{}) {
	if logLevel > 3 {
		return
	}
	write(warnPrefix, o)
}

func Error(o ...interface{}) {
	if logLevel > 4 {
		return
	}
	write(errorPrefix, o)
}

// Fatal this will cause system exists!!!
func Fatal(o ...interface{}) {
	if logLevel > 5 {
		return
	}
	write(fatalPrefix, o)
	os.Exit(0)
}

func write(levelPrefix string, o ...interface{}) {
	line := fmt.Sprint(o)
	ts := timeutil.GetLongLongDateString(time.Now())
	var buff bytes.Buffer
	buff.WriteString(levelPrefix)
	buff.WriteString(ts)
	buff.WriteString(" ")
	getCaller(&buff)
	buff.WriteString(" ")
	buff.WriteString(strings.TrimRight(strings.TrimLeft(line, "["), "]"))
	// buff.WriteString(line)
	buff.WriteString("\n")

	fmt.Print(string(buff.Bytes()))
	if out != nil {
		check()
		out.Write(buff.Bytes())
	}
}

func getCaller(buff *bytes.Buffer) {
	_, file, line, success := runtime.Caller(3)
	if success {
		buff.WriteString("[")
		buff.WriteString(file[strings.LastIndex(file, "/")+1:])
		buff.WriteString(":")
		buff.WriteString(strconv.Itoa(line))
		buff.WriteString("]")
	}
}
