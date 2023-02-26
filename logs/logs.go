package logs

import (
	"fmt"
	"io"
	"log"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
)

type logWriter struct {
	writer io.Writer
	enable bool
	level  string
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	if !writer.enable {
		return 0, nil
	}
	message := []byte(fmt.Sprintf("[%s][%s] %s", time.Now().Format(time.RFC3339), writer.level, bytes))
	return writer.writer.Write(message)
}

func InitLoggers(output io.Writer, info, warn, debug bool) {
	WarningLogger = log.New(logWriter{writer: output, enable: warn, level: "WARN"}, "", log.Lshortfile)
	InfoLogger = log.New(logWriter{writer: output, enable: info, level: "INFO"}, "", log.Lshortfile)
	DebugLogger = log.New(logWriter{writer: output, enable: debug, level: "DEBUG"}, "", log.Lshortfile)
	ErrorLogger = log.New(logWriter{writer: output, enable: true, level: "ERROR"}, "", log.Lshortfile)
}
