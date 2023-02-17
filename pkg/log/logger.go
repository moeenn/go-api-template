package log

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Instance *log.Logger
}

func (l Logger) Log(level Level, message string) {
	line := fmt.Sprintf("[%s] %s", level.String(), message)
	l.Instance.Println(line)
}

func New() *Logger {
	return &Logger{
		Instance: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}
