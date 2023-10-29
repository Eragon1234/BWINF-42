package ansi

import (
	"fmt"
)

type Style string

func (s Style) Add(others ...Style) Style {
	if len(others) == 0 {
		return s
	}
	otherStyling := others[0].Add(others[1:]...)
	return Style(fmt.Sprintf("%s;%s", s, otherStyling))
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

func StyleString(s string, styles ...Style) string {
	if len(styles) == 0 {
		return s
	}
	styling := styles[0].Add(styles[1:]...)
	return fmt.Sprintf("\033[%sm%s%s", styling, s, Reset)
}

func S(s string, styles ...Style) string {
	return StyleString(s, styles...)
}

func RGB(r, g, b int) Style {
	return Style(fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
}
