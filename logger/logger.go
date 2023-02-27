package logger

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
	WarningLogger = log.New(logWriter{writer: output, enable: warn, level: "WARN"}, "", 0)
	InfoLogger = log.New(logWriter{writer: output, enable: info, level: "INFO"}, "", 0)
	DebugLogger = log.New(logWriter{writer: output, enable: debug, level: "DEBUG"}, "", log.Lshortfile)
	ErrorLogger = log.New(logWriter{writer: output, enable: true, level: "ERROR"}, "", log.Lshortfile)
}

func InfoPrint(v ...any) {
	InfoLogger.Print(v...)
}

func InfoPrintf(format string, v ...any) {
	InfoLogger.Printf(format, v...)
}

func InfoPrintln(v ...any) {
	InfoLogger.Println(v...)
}

func WarnPrint(v ...any) {
	WarningLogger.Print(v...)
}

func WarnPrintf(format string, v ...any) {
	WarningLogger.Printf(format, v...)
}

func WarnPrintln(v ...any) {
	WarningLogger.Println(v...)
}

func DebugPrint(v ...any) {
	DebugLogger.Print(v...)
}

func DebugPrintf(format string, v ...any) {
	DebugLogger.Printf(format, v...)
}

func DebugPrintln(v ...any) {
	DebugLogger.Println(v...)
}

func ErrorPrint(v ...any) {
	ErrorLogger.Print(v...)
}

func ErrorPrintf(format string, v ...any) {
	ErrorLogger.Printf(format, v...)
}

func ErrorPrintln(v ...any) {
	ErrorLogger.Println(v...)
}

func Fatal(v ...any) {
	ErrorLogger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	ErrorLogger.Fatalf(format, v...)
}

func Fatalln(v ...any) {
	ErrorLogger.Fatalln(v...)
}

func Panic(v ...any) {
	ErrorLogger.Panic(v...)
}

func Panicf(format string, v ...any) {
	ErrorLogger.Panicf(format, v...)
}

func Panicln(v ...any) {
	ErrorLogger.Panicln(v...)
}
