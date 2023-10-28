package ansi

import "fmt"

type Color string

const Reset = "\033[0m"

const (
	Black   Color = "\033[30m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"
)

func ColorString(s string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, s, Reset)
}

func C(s string, color Color) string {
	return ColorString(s, color)
}
