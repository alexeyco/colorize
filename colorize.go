package colorize

import "fmt"

const (
	// FgBlack black foreground
	FgBlack Style = iota + 30
	// FgRed red foreground
	FgRed
	// FgGreen green foreground
	FgGreen
	// FgYellow yellow foreground
	FgYellow
	// FgBlue blue foreground
	FgBlue
	// FgMagenta magenta foreground
	FgMagenta
	// FgCyan cyan foreground
	FgCyan
	// FgWhite white foreground
	FgWhite
	// FgDefault default foreground
	FgDefault Style = 39
)

const (
	// FgDarkGray dark gray foreground
	FgDarkGray Style = iota + 90
	// FgLightRed light light foreground
	FgLightRed
	// FgLightGreen light green foreground
	FgLightGreen
	// FgLightYellow light yellow foreground
	FgLightYellow
	// FgLightBlue light blue foreground
	FgLightBlue
	// FgLightMagenta light magenta foreground
	FgLightMagenta
	// FgLightCyan light cyan foreground
	FgLightCyan
	// FgLightWhite light white foreground
	FgLightWhite
)

const (
	// BgBlack black background
	BgBlack Style = iota + 40
	// BgRed red background
	BgRed
	// BgGreen green background
	BgGreen
	// BgYellow yellow background
	BgYellow
	// BgBlue blue background
	BgBlue
	// BgMagenta magenta background
	BgMagenta
	// BgCyan cyan background
	BgCyan
	// BgWhite white background
	BgWhite
	// BgDefault default background
	BgDefault Style = 49
)

const (
	// BgDarkGray dark gray background
	BgDarkGray Style = iota + 100
	// BgLightRed light red background
	BgLightRed
	// BgLightGreen light green background
	BgLightGreen
	// BgLightYellow light yellow background
	BgLightYellow
	// BgLightBlue light blue background
	BgLightBlue
	// BgLightMagenta light magenta background
	BgLightMagenta
	// BgLightCyan light cyan background
	BgLightCyan
	// BgLightWhite light white background
	BgLightWhite
)

const (
	// Reset reset all text styles to default
	Reset Style = iota
	// Bold bold text
	Bold
	// Fuzzy fuzzy text
	Fuzzy
	// Italic italic text
	Italic
	// Underscore underscored text
	Underscore
	// Blink blink text
	Blink
	// FastBlink fast blink text
	FastBlink
	// Reverse reverts text background and foregrond
	Reverse
	// Concealed concealed text
	Concealed
	// Strikethrough strike text
	Strikethrough
)

const tpl = "\x1b[%dm"

var reset = Reset.String()

// Style text style type
type Style int8

// String returns terminal style code
func (s Style) String() string {
	return fmt.Sprintf(tpl, s)
}

// Colorize applies styles to specified message
func Colorize(message interface{}, styles ...Style) string {
	var res string
	for _, style := range styles {
		res += style.String()
	}

	return res + fmt.Sprint(message) + reset
}
