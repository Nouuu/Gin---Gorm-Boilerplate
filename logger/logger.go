package logger

import (
	"fmt"
	"io"
	"log"
	"time"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	debugLogger   *log.Logger
	errorLogger   *log.Logger
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
	warningLogger = log.New(logWriter{writer: output, enable: warn, level: "WARN"}, "", 0)
	infoLogger = log.New(logWriter{writer: output, enable: info, level: "INFO"}, "", 0)
	debugLogger = log.New(logWriter{writer: output, enable: debug, level: "DEBUG"}, "", log.Lshortfile)
	errorLogger = log.New(logWriter{writer: output, enable: true, level: "ERROR"}, "", log.Lshortfile)
}

func InfoPrint(v ...any) {
	infoLogger.Print(v...)
}

func InfoPrintf(format string, v ...any) {
	infoLogger.Printf(format, v...)
}

func InfoPrintln(v ...any) {
	infoLogger.Println(v...)
}

func WarnPrint(v ...any) {
	warningLogger.Print(v...)
}

func WarnPrintf(format string, v ...any) {
	warningLogger.Printf(format, v...)
}

func WarnPrintln(v ...any) {
	warningLogger.Println(v...)
}

func DebugPrint(v ...any) {
	debugLogger.Print(v...)
}

func DebugPrintf(format string, v ...any) {
	debugLogger.Printf(format, v...)
}

func DebugPrintln(v ...any) {
	debugLogger.Println(v...)
}

func ErrorPrint(v ...any) {
	errorLogger.Print(v...)
}

func ErrorPrintf(format string, v ...any) {
	errorLogger.Printf(format, v...)
}

func ErrorPrintln(v ...any) {
	errorLogger.Println(v...)
}

func Fatal(v ...any) {
	errorLogger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	errorLogger.Fatalf(format, v...)
}

func Fatalln(v ...any) {
	errorLogger.Fatalln(v...)
}

func Panic(v ...any) {
	errorLogger.Panic(v...)
}

func Panicf(format string, v ...any) {
	errorLogger.Panicf(format, v...)
}

func Panicln(v ...any) {
	errorLogger.Panicln(v...)
}
