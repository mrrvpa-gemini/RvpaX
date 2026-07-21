package colors

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	White  = color.New(color.FgWhite).SprintFunc()
)

func Error(msg string) string {
	return Red("[!] " + msg)
}

func Success(msg string) string {
	return Green("[+] " + msg)
}

func Info(msg string) string {
	return Cyan("[*] " + msg)
}

func Warning(msg string) string {
	return Yellow("[!] " + msg)
}

func Bold(msg string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", msg)
}
