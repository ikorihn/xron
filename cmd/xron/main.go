package main

import (
	"strings"

	"github.com/ikorihn/xron"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<books>
    <book ID="extension1" available="yes">
        <title>Book Title 1</title>
        <price>100 - 200</price>
        <note name="mynote">
            <author id="20">Author 1</author>
        </note>
        <category name="novel" id="2" />
        <empty />
    </book>
    <book ID="Book001" available="no">
        <title>Book Title 2</title>
        <price>400 - 500</price>
    </book>
    <book ID="extension2">
        <title>Book Title 3</title>
        <price>abc - 500</price>
    </book>
    <book available="soon">
        <title>Book Title 5</title>
        <price>a100 - c200</price>
    </book>
    <book ID="Book003" available="yes">
        <title>Book Title 6</title>
        <price>c100 - 200</price>
    </book>
</books>
`

	xron.ConvertXmlToXpath(strings.NewReader(s))
}
