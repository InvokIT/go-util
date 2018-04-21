package debug

import (
	"log"
	"os"
	"fmt"
	"regexp"
)

var separatorRegex = regexp.MustCompile("[,;:\\s]")
var enabledPrefixRegexes []*regexp.Regexp

func init() {
	logger := log.New(os.Stderr, "github.com/invokit/vorspiel-lib/debug", log.LstdFlags)


	enabledPrefixesString := os.Getenv("DEBUG")
	enabledPrefixes := separatorRegex.Split(enabledPrefixesString, -1)

	for _, prefix := range enabledPrefixes {
		regex, err := regexp.Compile(fmt.Sprintf("^\\Q%s\\E(|$)", prefix))
		if err != nil {
			logger.Printf("Invalid value in DEBUG env-var: '%s'. Error message: %s", prefix, err)
			continue
		}

		enabledPrefixRegexes = append(enabledPrefixRegexes, regex)
	}
}

func NewLogger(packageName string) Logger {
	for _, prefixRegex := range enabledPrefixRegexes {
		if prefixRegex.MatchString(packageName) {
			logger := log.New(os.Stderr, fmt.Sprintf("DEBUG: %s", packageName), log.Llongfile | log.LstdFlags | log.Lmicroseconds)

			return &loggerImpl{logger:logger}
		}
	}

	return &noopLoggerInstance
}


type Logger interface {
	Print(v ...interface{})
	Printf(fmt string, values ...interface{})
	Println(v ...interface{})
}


type noopLogger struct {}

func (logger *noopLogger) Println(v ...interface{}) {}

func (logger *noopLogger) Print(v ...interface{}) {}

func (logger *noopLogger) Printf(fmt string, values ...interface{}) {}

var noopLoggerInstance = noopLogger{}


type loggerImpl struct {
	logger* log.Logger
}

func (l *loggerImpl) Print(v ...interface{}) {
	l.logger.Print(v)
}

func (l *loggerImpl) Println(v ...interface{}) {
	l.logger.Println(v)
}

func (l *loggerImpl) Printf(fmt string, values ...interface{}) {
	l.logger.Printf(fmt, values)
}
