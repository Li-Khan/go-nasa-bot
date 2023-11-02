package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	nameColor    = "\033[36m"
	resetColor   = "\033[0m"
	infoColor    = "\033[32m"
	warningColor = "\033[1;33m"
	debugColor   = "\033[34m"
	noticeColor  = "\u001B[96m"
	errorColor   = "\033[31m"
)

var (
	logPrefix = "%s %s:%s "
	info      = infoColor + "INFO:"
	warning   = warningColor + "WARNING:"
	debug     = debugColor + "DEBUG:"
	notice    = noticeColor + "NOTICE:"
	err       = errorColor + "ERROR:"
	name      = nameColor + "Unknown"

	infoPrefix    = prefix(info, name)
	warningPrefix = prefix(warning, name)
	debugPrefix   = prefix(debug, name)
	noticePrefix  = prefix(notice, name)
	errorPrefix   = prefix(err, name)

	Info       = logger(infoPrefix)
	Warning    = logger(warningPrefix)
	Debug      = logger(debugPrefix)
	Notice     = logger(noticePrefix)
	Error      = logger(errorPrefix)
	onceLogger sync.Once
)

func Init(n string) {
	onceLogger.Do(func() {
		n = strings.TrimSpace(n)
		if n == "" {
			n = name
		}
		n = nameColor + n

		infoPrefix = prefix(info, n)
		warningPrefix = prefix(warning, n)
		debugPrefix = prefix(debug, n)
		noticePrefix = prefix(notice, n)
		errorPrefix = prefix(err, n)

		Info = logger(infoPrefix)
		Warning = logger(warningPrefix)
		Debug = logger(debugPrefix)
		Notice = logger(noticePrefix)
		Error = logger(errorPrefix)
	})
}

func logger(prefix string) *log.Logger {
	return log.New(os.Stderr, prefix, log.Ltime|log.Lshortfile)
}

func prefix(log string, name string) string {
	return fmt.Sprintf(logPrefix, log, name, resetColor)
}
