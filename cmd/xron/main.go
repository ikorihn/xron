package main

import (
	"fmt"
	"os"

	"github.com/ikorihn/xron"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Fprintf(os.Stderr, `USAGE:
	xron < file.xml

Translates XML input to a greppable, simplified XPath-like line-by-line format.
`)
		os.Exit(1)
	}

	xron.ConvertXmlToXpathFunc(os.Stdin, func(row string) {
		fmt.Println(row)
	})
}
