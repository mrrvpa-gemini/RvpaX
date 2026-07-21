package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/mrrvpa-gemini/rvpa/pkg/colors"
)

type Logger struct {
	verbose bool
	file    *os.File
}

func New(verbose bool) *Logger {
	return &Logger{
		verbose: verbose,
	}
}

func (l *Logger) Info(msg string) {
	fmt.Println(colors.Info(msg))
}

func (l *Logger) Success(msg string) {
	fmt.Println(colors.Success(msg))
}

func (l *Logger) Warning(msg string) {
	fmt.Println(colors.Warning(msg))
}

func (l *Logger) Error(msg string) {
	fmt.Println(colors.Error(msg))
}

func (l *Logger) Debug(msg string) {
	if l.verbose {
		fmt.Println(colors.Blue("[DEBUG] " + msg))
	}
}

func (l *Logger) Fatal(msg string) {
	log.Fatal(colors.Error(msg))
}
