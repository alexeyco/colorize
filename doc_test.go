package colorize_test

import (
	"fmt"

	"github.com/alexeyco/colorize"
)

// ExampleColorize is a simple example
func ExampleColorize() {
	message := "How much wood could a woodchuck chuck if a woodchuck could chuck wood?"
	fmt.Println(colorize.Colorize(message, colorize.Bold, colorize.BgGreen, colorize.FgRed))
}
