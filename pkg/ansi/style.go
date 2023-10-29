package ansi

import (
	"fmt"
	"strings"
)

type Style string

func (s Style) Add(others ...Style) Style {
	var sb strings.Builder
	sb.WriteString(string(s))
	for _, other := range others {
		sb.WriteRune(';')
		sb.WriteString(string(other))
	}
	return Style(fmt.Sprintf("%s;%s", s, sb.String()))
}

const Reset = "\033[0m"

const (
	Black         Style = "30"
	Red           Style = "31"
	Green         Style = "32"
	Yellow        Style = "33"
	Blue          Style = "34"
	Magenta       Style = "35"
	Cyan          Style = "36"
	White         Style = "37"
	BrightBlack   Style = "90"
	BrightRed     Style = "91"
	BrightGreen   Style = "92"
	BrightYellow  Style = "93"
	BrightBlue    Style = "94"
	BrightMagenta Style = "95"
	BrightCyan    Style = "96"
	BrightWhite   Style = "97"
	Bold          Style = "1"
	Underline     Style = "4"
)

const emptyStyle = Style("")

func StyleString(s string, styles ...Style) string {
	styling := emptyStyle.Add(styles...)
	return fmt.Sprintf("\033[%sm%s%s", styling, s, Reset)
}

func S(s string, styles ...Style) string {
	return StyleString(s, styles...)
}

func RGB(r, g, b int) Style {
	return Style(fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
}
