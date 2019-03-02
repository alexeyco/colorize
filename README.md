# Colorize

[![Travis](https://img.shields.io/travis/alexeyco/colorize.svg)](https://travis-ci.org/alexeyco/colorize)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/colorize/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/colorize?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/colorize)](https://goreportcard.com/report/github.com/alexeyco/colorize)
[![GoDoc](https://godoc.org/github.com/alexeyco/colorize?status.svg)](https://godoc.org/github.com/alexeyco/colorize)
[![license](https://img.shields.io/github/license/alexeyco/colorize.svg)](https://github.com/alexeyco/colorize)

Text styles for terminal

![Example](https://raw.githubusercontent.com/alexeyco/colorize/master/example.png)

## Example

```go
package main

import (
	"fmt"

	"github.com/alexeyco/colorize"
)

func main() {
	message := "How much wood could a woodchuck chuck if a woodchuck could chuck wood?"
	fmt.Println(colorize.Colorize(message, colorize.Bold, colorize.BgGreen, colorize.FgRed))
}
```

## Inspired by (see also)

* [gookit/color](https://github.com/gookit/color)
* [logrusorgru/aurora](https://github.com/logrusorgru/aurora)

## License

```
MIT License

Copyright (c) 2019 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
