package colorize_test

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/alexeyco/colorize"
)

var (
	styles = []colorize.Style{
		colorize.FgBlack,
		colorize.FgRed,
		colorize.FgGreen,
		colorize.FgYellow,
		colorize.FgBlue,
		colorize.FgMagenta,
		colorize.FgCyan,
		colorize.FgWhite,
		colorize.FgDefault,
		colorize.FgDarkGray,
		colorize.FgLightRed,
		colorize.FgLightGreen,
		colorize.FgLightYellow,
		colorize.FgLightBlue,
		colorize.FgLightMagenta,
		colorize.FgLightCyan,
		colorize.FgLightWhite,
		colorize.BgBlack,
		colorize.BgRed,
		colorize.BgGreen,
		colorize.BgYellow,
		colorize.BgBlue,
		colorize.BgMagenta,
		colorize.BgCyan,
		colorize.BgWhite,
		colorize.BgDefault,
		colorize.BgDarkGray,
		colorize.BgLightRed,
		colorize.BgLightGreen,
		colorize.BgLightYellow,
		colorize.BgLightBlue,
		colorize.BgLightMagenta,
		colorize.BgLightCyan,
		colorize.BgLightWhite,
		colorize.Bold,
		colorize.Fuzzy,
		colorize.Italic,
		colorize.Underscore,
		colorize.Blink,
		colorize.FastBlink,
		colorize.Reverse,
		colorize.Concealed,
		colorize.Strikethrough,
	}

	message = "How much wood could a woodchuck chuck if a woodchuck could chuck wood?"
)

// TestColorize test colorize
func TestColorizeSimple(t *testing.T) {
	for _, style := range styles {
		result := colorize.Colorize(message, style)
		if err := hasStyles(result, style); err != nil {
			t.Error(err)
		}
	}
}

// TestColorize composite style
func TestColorizeComposite(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		s := randomStyle(rand.Intn(3) + 2)
		result := colorize.Colorize(message, s...)

		if err := hasStyles(result, s...); err != nil {
			t.Error(err)
		}
	}
}

func hasStyles(message string, styles ...colorize.Style) error {
	styles = append(styles, colorize.Reset)
	for _, style := range styles {
		if !strings.Contains(message, style.String()) {
			return errors.New(fmt.Sprintf("must contain %d style", style))
		}
	}

	return nil
}

func randomStyle(count int, exclude ...colorize.Style) []colorize.Style {
	var s []colorize.Style

	for len(s) < count {
		style := styles[rand.Intn(len(styles))]

		if !styleExist(style, s) {
			s = append(s, style)
		}
	}

	return s
}

func styleExist(style colorize.Style, slice []colorize.Style) bool {
	for _, s := range slice {
		if s == style {
			return true
		}
	}

	return false
}
