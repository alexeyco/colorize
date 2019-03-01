package colorize

import "fmt"

const (
	FgBlack Style = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// FgDefault revert default FG
	FgDefault Style = 39
)

const (
	FgDarkGray Style = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
)

const (
	BgBlack Style = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Style = 49
)

const (
	BgDarkGray Style = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
)

// Option settings
const (
	Reset Style = iota
	Bold
	Fuzzy
	Italic
	Underscore
	Blink
	FastBlink
	Reverse
	Concealed
	Strikethrough
)

const (
	tpl   = "\x1b[%dm"
)

type Style int8

func (s Style) String() string {
	return fmt.Sprintf(tpl, s)
}

func Colorize(message interface{}, styles ...Style) string {
	var res string
	for _, style := range styles {
		res += style.String()
	}

	return res + fmt.Sprint(message) + Reset.String()
}
