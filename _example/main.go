package main

import (
	"fmt"
	"github.com/alexeyco/colorize"
)

type Column struct {
	Text string
	Style colorize.Style
}

type Row []Column

var rows = []Row{
	{
		Column{
			Text: "Black foreground",
			Style: colorize.FgBlack,
		},
		Column{
			Text: "Black background",
			Style: colorize.BgBlack,
		},
		Column{
			Text: "Dark gray foreground",
			Style: colorize.FgDarkGray,
		},
		Column{
			Text: "Dark gray background",
			Style: colorize.BgDarkGray,
		},
	},
	{
		Column{
			Text: "Red foreground",
			Style: colorize.FgRed,
		},
		Column{
			Text: "Red background",
			Style: colorize.BgRed,
		},
		Column{
			Text: "Light red foreground",
			Style: colorize.FgLightRed,
		},
		Column{
			Text: "Light red background",
			Style: colorize.BgLightRed,
		},
	},
	{
		Column{
			Text: "Green foreground",
			Style: colorize.FgGreen,
		},
		Column{
			Text: "Green background",
			Style: colorize.BgGreen,
		},
		Column{
			Text: "Light green foreground",
			Style: colorize.FgLightGreen,
		},
		Column{
			Text: "Light green background",
			Style: colorize.BgLightGreen,
		},
	},
	{
		Column{
			Text: "Yellow foreground",
			Style: colorize.FgYellow,
		},
		Column{
			Text: "Yellow background",
			Style: colorize.BgYellow,
		},
		Column{
			Text: "Light yellow foreground",
			Style: colorize.FgLightYellow,
		},
		Column{
			Text: "Light yellow background",
			Style: colorize.BgLightYellow,
		},
	},
	{
		Column{
			Text: "Blue foreground",
			Style: colorize.FgBlue,
		},
		Column{
			Text: "Blue background",
			Style: colorize.BgBlue,
		},
		Column{
			Text: "Light blue foreground",
			Style: colorize.FgLightBlue,
		},
		Column{
			Text: "Light blue background",
			Style: colorize.BgLightBlue,
		},
	},
	{
		Column{
			Text: "Magenta foreground",
			Style: colorize.FgMagenta,
		},
		Column{
			Text: "Magenta background",
			Style: colorize.BgMagenta,
		},
		Column{
			Text: "Light magenta foreground",
			Style: colorize.FgLightMagenta,
		},
		Column{
			Text: "Light magenta background",
			Style: colorize.BgLightMagenta,
		},
	},
	{
		Column{
			Text: "Cyan foreground",
			Style: colorize.FgCyan,
		},
		Column{
			Text: "Cyan background",
			Style: colorize.BgCyan,
		},
		Column{
			Text: "Light cyan foreground",
			Style: colorize.FgLightCyan,
		},
		Column{
			Text: "Light cyan background",
			Style: colorize.BgLightCyan,
		},
	},
	{
		Column{
			Text: "White foreground",
			Style: colorize.FgWhite,
		},
		Column{
			Text: "White background",
			Style: colorize.BgWhite,
		},
		Column{
			Text: "Light white foreground",
			Style: colorize.FgLightWhite,
		},
		Column{
			Text: "Light white background",
			Style: colorize.BgLightWhite,
		},
	},
	{
		Column{
			Text: "Bold text",
			Style: colorize.Bold,
		},
		Column{
			Text: "Fuzzy text",
			Style: colorize.Fuzzy,
		},
		Column{
			Text: "Strike text",
			Style: colorize.Strikethrough,
		},
		Column{
			Text: "Italic text",
			Style: colorize.Italic,
		},
	},
	{
		Column{
			Text: "Underscore text",
			Style: colorize.Underscore,
		},
		Column{
			Text: "Blink text",
			Style: colorize.Blink,
		},
		Column{
			Text: "Reverse text",
			Style: colorize.Reverse,
		},
	},
}

func main()  {
	for _, row := range rows {
		for _, column := range row {
			fmt.Print(colorize.Colorize(fmt.Sprintf("%-24v", column.Text), column.Style))
			fmt.Print("  ")
		}

		fmt.Println()
	}
}
